package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type CertSummary struct {
	Digest    string `json:"digest"`
	Algorithm string `json:"algorithm"`
	Status    bool   `json:"status"`
}

type CertInfo struct {
	CertSummary CertSummary `json:"certSummary"`
	CertInPEM   string      `json:"CertInPEM"`
}

func main() {

	digest := "adjfasfk"
	algorithm := "SHA256"
	status := true
	var certSummary = CertSummary{Digest: digest, Algorithm: algorithm, Status: status}
	fmt.Println(certSummary)
	certInPEM := "AFJFASDFJ"
	var certInfo = CertInfo{CertSummary: certSummary, CertInPEM: certInPEM}
	fmt.Println(certInfo)
	jsonCertInfo, _ := json.Marshal(certInfo)
	fmt.Println("certInfo in bytes:", jsonCertInfo)
	fmt.Println("certInfo in json:", string(jsonCertInfo))
	hashcode := GetSHA256HashCode(jsonCertInfo)
	fmt.Println("hashcode:", hashcode)
}

//Generate the hash vaule using SHA256
func GetSHA256HashCode(message []byte) string {
	bytes := sha256.Sum256(message)
	hashcode := hex.EncodeToString(bytes[:])
	return hashcode
}
