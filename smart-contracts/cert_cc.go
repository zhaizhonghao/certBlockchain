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
type CertSummary struct {
	digest    string  `json:"digest"`
	algorithm string  `json:"algorithm"`
	status    boolean `json:"status"`
}

// OwnerPrefix is used for creating the key for balances
const OwnerPrefix = "owner."

// Init Implements the Init method
// Receives 4 parameters =  [0] Symbol [1] TotalSupply   [2] Description  [3] Owner
func (certChain *CertChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")
	_, args := stub.GetFunctionAndParameters()
	// Check if we received the right number of arguments
	if len(args) < 4 {
		return shim.Error("Failed - incorrect number of parameters!!???")
	}
	symbol := string(args[0])
	// Get total supply & check if it is > 0
	totalSupply, err := strconv.ParseUint(string(args[1]), 10, 64)

	if err != nil || totalSupply == 0 {
		return shim.Error("Total Supply MUST be a number > 0 !! ")
	}

	// Creator name cannot be zeo length
	if len(args[3]) == 0 {
		return errorResponse("Creator identity cannot be 0 length!!!", 3)
	}
	creator := string(args[3])

	// Create an instance of the token struct
	var erc20 = ERC20Token{Symbol: symbol, TotalSupply: totalSupply, Description: string(args[2]), Creator: creator}

	// Convert to JSON and store token in the state
	jsonERC20, _ := json.Marshal(erc20)
	stub.PutState("token", []byte(jsonERC20))

	// Maintain the balances in the state db
	// In the begining all tokens are owned by the creator of the token
	key := OwnerPrefix + creator
	fmt.Println("Key=", key)
	err = stub.PutState(key, []byte(args[1]))
	if err != nil {
		return errorResponse(err.Error(), 4)
	}

	return shim.Success([]byte(jsonERC20))
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
 **/
func addCert(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//Check if enough arguments are in the args
	if len(args) < 5 {
		return errorResponse("Needs domainName, digest, algorithm & status!!!", 700)
	}

	//Check whether the certificate corresponding to the domain has existed
	bytes, _ := stub.GetState(domainName)
	if len(bytes) != 0 {
		// That means 0 token balance
		return errorResponse("The certificate corresponding to the domain has already existed", 703)
	}

	//Check whether the certificate in question has been expired
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		return errorResponse("failed to parse certificate PEM!!!", 700)
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return errorResponse("failed to parse certificate: "+err.Error(), 700)
	}
	now := time.Now()
	if now.Sub(notAfter) > 0 {
		return errorResponse("The certificate in question has expired!!!", 700)
	}

	//Check the digest
	algorithm := string(args[2])
	if algorithm != "SHA256" {
		return errorResponse("The algorithm of hashcode needs to be SHA256!!!", 700)
	}
	hashCode := GetSHA256HashCode([]byte(args[4]))
	digest := string(args[1])
	if digest != hashcode {
		return errorResponse("The digest is inconsistent with the hash of the certificate!!!", 700)
	}

	//Add the certificate

}

/**
 * Setter function
 * function deleteCert(address certOwnerID) public view returns (bool success);
 * Returns bool success OR failure
 * {"Args":["deleteCert","certOwnerID"]}
 **/
func deleteCert(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

/**
 * Setter function
 * function modifyCert(address certOwnerID, CertSummary certSummary) public view returns (bool success);
 * Returns bool success OR failure
 * {"Args":["modifyCert","certOwnerID","digest","algorithm","status"]}
 **/
func modifyCert(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

/**
 * Getter function
 * function getCert(address certOwnerID) public view returns (CertSummary certSummary);
 * Returns the summary of the certificate
 * {"Args":["getCert","certOwnerID"]}
 **/
func getCert(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

/**
 * Getter function
 * function totalSupply() public view returns (bool success);
 * Returns the totalSupply for the token
 **/
func totalSupply(stub shim.ChaincodeStubInterface) peer.Response {

	bytes, err := stub.GetState("token")
	if err != nil {
		return errorResponse(err.Error(), 5)
	}

	// Read the JSON and Initialize the struct
	var erc20 ERC20Token
	_ = json.Unmarshal(bytes, &erc20)

	// Create the JSON Response with totalSupply
	return successResponse(strconv.FormatUint(erc20.TotalSupply, 10))
}

/**
 * Getter function
 * function balanceOf(address tokenOwner) public view returns (uint balance);
 * Returns the balance for the specified owner
 * {"Args":["balanceOf","owner-id"]}
 **/
func balanceOf(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// Check if owner id is in the arguments
	if len(args) < 1 {
		return errorResponse("Needs OwnerID!!!", 6)
	}
	OwnerID := args[0]
	bytes, err := stub.GetState(OwnerPrefix + OwnerID)
	if err != nil {
		return errorResponse(err.Error(), 7)
	}

	response := balanceJSON(OwnerID, string(bytes))

	return successResponse(response)
}

/**
 * Setter function
 * function transfer(address to, uint tokens) public returns (bool success);
 * Transfer tokens
 * {"Args":["from","to","amount"]}
 **/
func transfer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// Check if owner id is in the arguments
	if len(args) < 3 {
		return errorResponse("Needs to, from & amount!!!", 700)
	}

	from := string(args[0])
	to := string(args[1])
	amount, err := strconv.Atoi(string(args[2]))
	if err != nil {
		return errorResponse(err.Error(), 701)
	}
	if amount <= 0 {
		return errorResponse("Amount MUST be > 0!!!", 702)
	}

	// Get the Balance for from
	bytes, _ := stub.GetState(OwnerPrefix + from)
	if len(bytes) == 0 {
		// That means 0 token balance
		return errorResponse("Balance MUST be > 0!!!", 703)
	}

	fromBalance, _ := strconv.Atoi(string(bytes))
	if fromBalance < amount {
		return errorResponse("Insufficient balance to cover transfer!!!", 704)
	}
	// Reduce the tokens in from account
	fromBalance = fromBalance - amount

	// Get the balance in to account
	bytes, _ = stub.GetState(OwnerPrefix + to)
	toBalance := 0
	if len(bytes) > 0 {
		toBalance, _ = strconv.Atoi(string(bytes))
	}
	toBalance += amount

	// Update the balance
	bytes = []byte(strconv.FormatInt(int64(fromBalance), 10))
	err = stub.PutState(OwnerPrefix+from, bytes)

	bytes = []byte(strconv.FormatInt(int64(toBalance), 10))
	err = stub.PutState(OwnerPrefix+to, bytes)

	// Emit Transfer Event
	eventPayload := "{\"from\":\"" + from + "\", \"to\":\"" + to + "\",\"amount\":" + strconv.FormatInt(int64(amount), 10) + "}"
	stub.SetEvent("transfer", []byte(eventPayload))
	return successResponse("\"Transfer Successful!!!\"")
}

// balanceJSON creates a JSON for representing the balance
func balanceJSON(OwnerID, balance string) string {
	return "{\"owner\":\"" + OwnerID + "\", \"balance\":" + balance + "}"
}

func getHistoryByKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// Check if owner id is in the arguments
	if len(args) < 1 {
		return errorResponse("Needs OwnerID!!!", 6)
	}
	//To create the key
	key := OwnerPrefix + args[0]
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
