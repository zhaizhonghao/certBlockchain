package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// CertSummary structure manages the state
type CertSummary struct {
	Status     bool   `json:"status"`
	Algorithm  string `json:"algorithm"`
	DomainName string `json:"domainName"`
	Digest     string `json:"digest"`
}

type CertInfo struct {
	Status     bool   `json:"status"`
	Algorithm  string `json:"algorithm"`
	DomainName string `json:"domainName"`
	CertInPEM  string `json:"CertInPEM"`
}

func main() {

	const certPEM = `-----BEGIN CERTIFICATE-----
MIIGcjCCBVqgAwIBAgIMBgsWXop3UaR6Fo1hMA0GCSqGSIb3DQEBCwUAMGYxCzAJ
BgNVBAYTAkJFMRkwFwYDVQQKExBHbG9iYWxTaWduIG52LXNhMTwwOgYDVQQDEzNH
bG9iYWxTaWduIE9yZ2FuaXphdGlvbiBWYWxpZGF0aW9uIENBIC0gU0hBMjU2IC0g
RzIwHhcNMTkwNTA2MTI0OTI2WhcNMjAwOTIzMjI1OTM4WjBmMQswCQYDVQQGEwJV
UzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEU
MBIGA1UEChMLVWRlbXksIEluYy4xFDASBgNVBAMMCyoudWRlbXkuY29tMIIBIjAN
BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtmUC3p/7EpypE+Dz+gNClvpiw/my
FFNwBlVRvCZ7ZEGXXLIg2XQnJdjAcc8dGgJcivDvVEWLSzXZyi4tBL02deIXDYRm
HrUrSCLB449EjjAeCrqhC904BIUDbVV6+caK5y0gTsSidetCAy1G1BMt6INDdjBR
fsYntvhLGMLL5uyZ38dinGfwmBz+nisgQ1TcQNaec5gsLb0KuO/RB3+bciOMDghh
jkSyfacEaiYfdIB3NnWu9uXK/C68KN6440q9gujv9oQ6fFrGhiu1bTlT8sTmjV5o
0SuzI9Q8xpxJ4Iiq4L+4JUWDtNRLLx+V2RSTrNgGPLfqLEcX02cuPdZIuwIDAQAB
o4IDHjCCAxowDgYDVR0PAQH/BAQDAgWgMIGgBggrBgEFBQcBAQSBkzCBkDBNBggr
BgEFBQcwAoZBaHR0cDovL3NlY3VyZS5nbG9iYWxzaWduLmNvbS9jYWNlcnQvZ3Nv
cmdhbml6YXRpb252YWxzaGEyZzJyMS5jcnQwPwYIKwYBBQUHMAGGM2h0dHA6Ly9v
Y3NwMi5nbG9iYWxzaWduLmNvbS9nc29yZ2FuaXphdGlvbnZhbHNoYTJnMjBWBgNV
HSAETzBNMEEGCSsGAQQBoDIBFDA0MDIGCCsGAQUFBwIBFiZodHRwczovL3d3dy5n
bG9iYWxzaWduLmNvbS9yZXBvc2l0b3J5LzAIBgZngQwBAgIwCQYDVR0TBAIwADAh
BgNVHREEGjAYggsqLnVkZW15LmNvbYIJdWRlbXkuY29tMB0GA1UdJQQWMBQGCCsG
AQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQU0bmIONs2Ecw4IH2t0IcuXW5bJK0w
HwYDVR0jBBgwFoAUlt5h8b0cFilTHMDMfTuDAEDmGnwwggF+BgorBgEEAdZ5AgQC
BIIBbgSCAWoBaAB3AFWB1MIWkDYBSuoLm1c8U/DA5Dh4cCUIFy+jqh0HE9MMAAAB
ao0xIY8AAAQDAEgwRgIhAPkfrrD0CxdAwswgWf6jhT7dd+ta1U9o52bH6GMrpFpP
AiEA31wxSDDCZIGAg+x5qbOOwkCuw3RzfxN+VpLEAkfd9S8AdgCkuQmQtBhYFIe7
E6LMZ3AKPDWYBPkb37jjd80OyA3cEAAAAWqNMSF2AAAEAwBHMEUCIFE5yJaIfeyI
+RTY22NqhOevkAANFV58muomn6BRhAzPAiEA0XqEVHeupxab7vYFjWQvOIKtib1e
QW7eUqmmJ/0ltywAdQBvU3asMfAxGdiZAKRRFf93FRwR2QLBACkGjbIImjfZEwAA
AWqNMSHHAAAEAwBGMEQCIEdNV0ULjMdB9ssjYaX2uRdgNE4x3TWCFVgUAj3XA3qT
AiA6VXlQ8gRTbe4GCc/EqF6mBEASHi34zDTTNws440gMzjANBgkqhkiG9w0BAQsF
AAOCAQEAGbwsHU9PlTagyC/eH7G5XhuzcDbkTLUg7e2nvy6a1dGtDXhmV19NoGc/
uwph/E1KTsr9gtp1IDl1U+tgEt62YnLlvJM1lq4mQ5UVUUW+AMI8+PNs/4Mh1N5s
LXPWPoqpp6vvgo9+zZ5Si6rYwCdILqh5lhv4FRwKQOYInbfIazgOoIwEIKa3DBYK
XMkE/+JgkXzAeZXyZroZMEzmEHA0NcGy6Yw9CY3/xsJsdlS2dNXJgIMUMZ90Fvkg
D1QapIZqWwlGLdV+ErUyJpycw6umzd0059f5vLblSOizPLJT5086GX1r/LaTE5ED
t0SjiH8QPqb5z5KRspazvD9ijH3+BA==
-----END CERTIFICATE-----`
	status := true
	algorithm := "SHA256"
	domainName := "zhaizhonghao.com"

	var certInfo = CertInfo{Status: status, Algorithm: algorithm, DomainName: domainName, CertInPEM: certPEM}
	jsonCertInfo, _ := json.Marshal(certInfo)
	digest := GetSHA256HashCode(jsonCertInfo)
	fmt.Println("digest:", digest)

}

//Generate the hash vaule using SHA256
func GetSHA256HashCode(message []byte) string {
	bytes := sha256.Sum256(message)
	hashcode := hex.EncodeToString(bytes[:])
	return hashcode
}
