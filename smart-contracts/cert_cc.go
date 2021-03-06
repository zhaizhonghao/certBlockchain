/**
 * This implements
 *
 **/
package main

import (
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"time"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"

	// Conversion functions
	"strconv"

	// JSON Encoding
	"encoding/hex"
	"encoding/json"
	"encoding/pem"

	// KV Interface
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
)

// CertChaincode Represents our chaincode object
type CertChaincode struct {
}

// CertSummary structure manages the state
//The issuer is the hash of the invoker's identity
type CertSummary struct {
	Status     bool   `json:"status"`
	Algorithm  string `json:"algorithm"`
	DomainName string `json:"domainName"`
	Issuer     string `json:"Issuer"`
	Digest     string `json:"digest"`
}

//The issuer is the hash of the invoker's identity
type CertInfo struct {
	Status     bool   `json:"status"`
	Algorithm  string `json:"algorithm"`
	DomainName string `json:"domainName"`
	Issuer     string `json:"Issuer"`
	CertInPEM  string `json:"CertInPEM"`
}

type CertBlockchain struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

// Init Implements the Init method
// Receives 2 parameters =  [0] Symbol  [1] Description
func (certChain *CertChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	//TODO
	// Simply print a message
	fmt.Println("Init executed")
	_, args := stub.GetFunctionAndParameters()
	// Check if we received the right number of arguments
	if len(args) < 2 {
		return shim.Error("Failed - incorrect number of parameters!!???")
	}
	symbol := string(args[0])
	description := string(args[1])
	//Create an instance of the CertChain struct
	var certBlockchain = CertBlockchain{Symbol: symbol, Description: description}
	//Convert to JSON and store the certChain in the state
	jsonCertBlockchain, _ := json.Marshal(certBlockchain)
	err := stub.PutState("certBlockchain", []byte(jsonCertBlockchain))
	if err != nil {
		return errorResponse(err.Error(), 4)
	}
	return shim.Success([]byte(jsonCertBlockchain))
}

// Invoke method
func (certChain *CertChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the function name and parameters
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Invoke executed : ", function, " args=", args)

	switch {

	// Query function
	case function == "addCert":
		return addCert(stub, args)
	case function == "deleteCert":
		return deleteCert(stub, args)
	case function == "modifyCert":
		return modifyCert(stub, args)
	case function == "getCert":
		return getCert(stub, args)
	case function == "getHistoryByKey":
		fmt.Println("enter the getHistoryByKey")
		return getHistoryByKey(stub, args)
	}

	return errorResponse("Invalid function", 1)
}

/**
 * Setter function
 * function addCert(address certOwnerID, CertSummary certSummary) public view returns (bool success);
 * Returns bool success OR failure
 * {"Args":["addCert","domainName","digest","algorithm","status","certInPem"]}
 * the digest is the hash value of {{}}
 **/
func addCert(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//Check if enough arguments are in the args
	if len(args) < 5 {
		return errorResponse("Needs domainName, digest, algorithm , status  and certInPem!!!", 700)
	}

	//Check whether the certificate corresponding to the domain has existed and not been revoked)
	domainName := string(args[0])
	bytes, _ := stub.GetState(domainName)
	if len(bytes) != 0 {
		tempCertSummary := CertSummary{}
		errGet := json.Unmarshal(bytes, &tempCertSummary)
		if errGet != nil {
			return errorResponse("Unkown expection on parsing the certSummary", 703)
		}
		if !tempCertSummary.Status {
			// That means the certificate has existed and not been revoked
			return errorResponse("The certificate corresponding to the domain has already existed and not been revoked", 703)
		}
	}

	//Check whether the certificate in question has been expired
	block, _ := pem.Decode([]byte(string(args[4])))
	if block == nil {
		return errorResponse("failed to parse certificate PEM!!!", 700)
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return errorResponse("failed to parse certificate: "+err.Error(), 700)
	}
	now := time.Now()
	if now.Sub(cert.NotAfter) > 0 {
		return errorResponse("The certificate in question has expired!!!", 700)
	}

	//Check the digest
	algorithm := string(args[2])
	if algorithm != "SHA256" {
		return errorResponse("The algorithm of hashcode needs to be SHA256!!!", 700)
	}
	status, err := strconv.ParseBool(args[3])
	if err != nil {
		return errorResponse("The status must be TRUE or FALSE!!!", 700)
	}
	//To get the GetCreator()
	creator, err := stub.GetCreator()
	if err != nil {
		return errorResponse("Failed to get the creator!!!", 6)
	}
	issuer := GetSHA256HashCode(creator)
	fmt.Println("issuer:", issuer)
	var certInfo = CertInfo{Status: status, Algorithm: algorithm, DomainName: domainName, Issuer: issuer, CertInPEM: string(args[4])}
	jsonCertInfo, _ := json.Marshal(certInfo)
	hashCode := GetSHA256HashCode(jsonCertInfo)
	digest := string(args[1])
	if digest != hashCode {
		return errorResponse("The digest is inconsistent with the hash of the certificate!!!", 700)
	}

	//Add the certificate
	var certSummary = CertSummary{Status: status, Algorithm: algorithm, DomainName: domainName, Issuer: issuer, Digest: digest}
	// Convert to JSON and store certSummary in the state
	jsonCertSummary, _ := json.Marshal(certSummary)
	err = stub.PutState(domainName, []byte(jsonCertSummary))
	if err != nil {
		return errorResponse(err.Error(), 4)
	}
	return successResponse("\"Add Successfully!!!\"")
}

/**
 * Setter function
 * function deleteCert(address certOwnerID) public view returns (bool success);
 * Returns bool success OR failure
 * {"Args":["deleteCert","certOwnerID"]}
 **/
func deleteCert(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// Check if owner id is in the arguments
	if len(args) < 1 {
		return errorResponse("Needs certificate OwnerID!!!", 6)
	}
	certOwnerID := args[0]

	//Only the issuer of the certificate can delete the certificate
	creator, err := stub.GetCreator()
	if err != nil {
		return errorResponse("Failed to get the creator!!!", 6)
	}
	hashOfCert := GetSHA256HashCode(creator)
	domainName := string(args[0])
	bytes, _ := stub.GetState(domainName)
	if len(bytes) == 0 {
		// That means the certificate doesn't exist
		return errorResponse("The certificate corresponding to the domain doesn't exist", 703)
	}
	tempCertSummary := CertSummary{}
	errGet := json.Unmarshal(bytes, &tempCertSummary)
	if errGet != nil {
		return errorResponse("Unkown expection on parsing the certSummary", 703)
	}
	if tempCertSummary.Issuer != hashOfCert {
		return errorResponse("Fail to delete, permission deny!!!", 703)
	}

	err = stub.DelState(certOwnerID)
	if err != nil {
		return errorResponse(err.Error(), 7)
	}
	return successResponse("\"Delete Successful!!!\"")
}

/**
 * Setter function
 * function modifyCert(address certOwnerID, CertSummary certSummary) public view returns (bool success);
 * Returns bool success OR failure
 * {"Args":["modifyCert","domainName","digest","algorithm","status","certInPem"]}
 **/
func modifyCert(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//Check if enough arguments are in the args
	if len(args) < 5 {
		return errorResponse("Needs domainName, digest, algorithm , status and certInPem!!!", 700)
	}

	//Only the issuer of the certificate can modify the certificate
	creator, err := stub.GetCreator()
	if err != nil {
		return errorResponse("Failed to get the creator!!!", 6)
	}
	hashOfCert := GetSHA256HashCode(creator)
	//check whether the certificate corresponding to the domain exists
	domainName := string(args[0])
	bytes, _ := stub.GetState(domainName)
	if len(bytes) == 0 {
		// That means the certificate doesn't exist
		return errorResponse("The certificate corresponding to the domain doesn't exist", 703)
	}
	tempCertSummary := CertSummary{}
	errGet := json.Unmarshal(bytes, &tempCertSummary)
	if errGet != nil {
		return errorResponse("Unkown expection on parsing the certSummary", 703)
	}
	if tempCertSummary.Issuer != hashOfCert {
		return errorResponse("Fail to modify, permission deny!!!", 703)
	}

	//Check the digest
	algorithm := string(args[2])
	if algorithm != "SHA256" {
		return errorResponse("The algorithm of hashcode needs to be SHA256!!!", 700)
	}
	status, err := strconv.ParseBool(args[3])
	if err != nil {
		return errorResponse("The status must be TRUE or FALSE!!!", 700)
	}
	var certInfo = CertInfo{Status: status, Algorithm: algorithm, DomainName: domainName, Issuer: hashOfCert, CertInPEM: string(args[4])}
	jsonCertInfo, _ := json.Marshal(certInfo)
	hashCode := GetSHA256HashCode(jsonCertInfo)
	digest := string(args[1])
	if digest != hashCode {
		return errorResponse("The digest is inconsistent with the hash of the certificate!!!", 700)
	}

	//Modify the certificate
	var certSummary = CertSummary{Status: status, Algorithm: algorithm, Issuer: hashOfCert, DomainName: domainName, Digest: digest}
	// Convert to JSON and store certSummary in the state
	jsonCertSummary, _ := json.Marshal(certSummary)
	err = stub.PutState(domainName, []byte(jsonCertSummary))
	if err != nil {
		return errorResponse(err.Error(), 4)
	}
	return successResponse("\"Modify Successfully!!!\"")
}

/**
 * Getter function
 * function getCert(address certOwnerID) public view returns (CertSummary certSummary);
 * Returns the summary of the certificate
 * {"Args":["getCert","certOwnerID"]}
 **/
func getCert(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	// Check if owner id is in the arguments
	if len(args) < 1 {
		return errorResponse("Needs certificate OwnerID!!!", 6)
	}
	certOwnerID := args[0]

	//Check if the certificate exists
	bytes, err := stub.GetState(certOwnerID)
	if err != nil {
		return errorResponse(err.Error(), 7)
	}
	if len(bytes) == 0 {
		// That means the certificate doesn't exist
		return errorResponse("The certificate corresponding to the domain doesn't exist", 703)
	}

	response := string(bytes)

	return successResponse(response)
}

/**
 * Getter function
 * function getHistoryByKey(address certOwnerID) public view returns (CertSummary certSummary);
 * Returns the summary of the certificate
 * {"Args":["getHistoryByKey","certOwnerID"]}
 **/
func getHistoryByKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// Check if owner id is in the arguments
	if len(args) < 1 {
		return errorResponse("Needs OwnerID!!!", 6)
	}
	//To create the key
	key := args[0]
	// Get the history for the key i.e., VIN#
	historyQueryIterator, err := stub.GetHistoryForKey(key)
	// In case of error - return error
	if err != nil {
		return shim.Error("Error in fetching history !!!" + err.Error())
	}
	// Local variable to hold the history record
	var resultModification *queryresult.KeyModification
	counter := 0
	resultJSON := "["
	// Start a loop with check for more rows
	for historyQueryIterator.HasNext() {

		// Get the next record
		resultModification, err = historyQueryIterator.Next()

		if err != nil {
			return shim.Error("Error in reading history record!!!" + err.Error())
		}

		// Append the data to local variable
		data := "{\"txn\":" + "\"" + resultModification.GetTxId() + "\""
		data += " , \"timestamp\": " + "\"" + resultModification.GetTimestamp().String() + "\""
		data += " , \"value\": " + string(resultModification.GetValue()) + "}  "
		if counter > 0 {
			data = ", " + data
		}
		resultJSON += data

		counter++
	}

	// Close the iterator
	historyQueryIterator.Close()

	// finalize the return string
	resultJSON += "]"
	resultJSON = "{ \"counter\": " + strconv.Itoa(counter) + ", \"txns\":" + resultJSON + "}"

	// return success
	return shim.Success([]byte(resultJSON))
}

//Generate the hash vaule using SHA256
func GetSHA256HashCode(message []byte) string {
	bytes := sha256.Sum256(message)
	hashcode := hex.EncodeToString(bytes[:])
	return hashcode
}

func errorResponse(err string, code uint) peer.Response {
	codeStr := strconv.FormatUint(uint64(code), 10)
	// errorString := "{\"error\": \"" + err +"\", \"code\":"+codeStr+" \" }"
	errorString := "{\"error\":" + err + ", \"code\":" + codeStr + " \" }"
	return shim.Error(errorString)
}

func successResponse(dat string) peer.Response {
	success := "{\"response\": " + dat + ", \"code\": 0 }"
	return shim.Success([]byte(success))
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Println("Started....")
	err := shim.Start(new(CertChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
