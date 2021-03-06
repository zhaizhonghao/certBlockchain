# Blockchain for certificate management
## Prerequisites

* Operating system : ubuntu >= 16.04 or centOS 7

* Memory: 8GB or above

* Golang (go1.14.6 linux/amd64 or above)

* Docker (following or above)

  ```
  Client:
   Version:           18.09.7
   API version:       1.39
   Go version:        go1.10.1
   Git commit:        2d0083d
   Built:             Fri Aug 16 14:20:06 2019
   OS/Arch:           linux/amd64
   Experimental:      false
  
  Server:
   Engine:
    Version:          18.09.7
    API version:      1.39 (minimum version 1.12)
    Go version:       go1.10.1
    Git commit:       2d0083d
    Built:            Wed Aug 14 19:41:23 2019
    OS/Arch:          linux/amd64
    Experimental:     false
  ```
  
* Docker-compose (version 1.17.1 or above)

* Pull some required docker images from the Docker Hub

  ```
  docker pull hyperledger/fabric-peer:latest
  docker pull hyperledger/fabric-orderer:latest
  docker pull hyperledger/fabric-tools:latest
  docker pull hyperledger/fabric-ccenv:latest
  docker pull hyperledger/fabric-baseos:latest
  docker pull hyperledger/fabric-kafka:latest
  docker pull hyperledger/fabric-zookeeper:latest
  docker pull hyperledger/fabric-couchdb:latest
  docker pull hyperledger/fabric-ca:latest
  ```

  

## Start a simple fabric network

* [Build Your First Network](https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html)

* Make a directory

  ```
  mkdir -p $GOPATH/src/github.com/hyperledger/fabric/singlepeer/chaincode/go/CertChain
  ```

* Copy the `cert_cc.go` and `govendor.sh` under the directory

* Install the  necessary dependencies by running

  ```
  . govendor.sh
  ```

## Installation & Initialization

```
docker exec -it cli bash
```

To install:
```
peer chaincode install -n certChain -p github.com/hyperledger/fabric/singlepeer/chaincode/go/CertChain/ -v 1.0
```
To instantiate:
```
peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n certChain -v 1.0 -c '{"Args":["init","CertChain","a blockchain for certificate management"]}' -P "AND ('Org1MSP.peer')"
```
## Smart Contract Operation Test
### addCert
```
peer chaincode invoke -C mychannel -n certChain -c '{"Args":["addCert","www.shenshimen.com","251fa59c6f6aee9b82d9d2378ce773632c28cbfd3265ec5c85538a43a41b998e","SHA256","false","-----BEGIN CERTIFICATE-----\nMIIGcjCCBVqgAwIBAgIMBgsWXop3UaR6Fo1hMA0GCSqGSIb3DQEBCwUAMGYxCzAJ\nBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9iYWxTaWduIG52LXNhMTwwOgYDVQQDEzNH\nbG9iYWxTaWduIE9yZ2FuaXphdGlvbiBWYWxpZGF0aW9uIENBIC0gU0hBMjU2IC0g\nRzIwHhcNMTkwNTA2MTI0OTI2WhcNMjAwOTIzMjI1OTM4WjBmMQswCQYDVQQGEwJV\nUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEU\nMBIGA1UEChMLVWRlbXksIEluYy4xFDASBgNVBAMMCyoudWRlbXkuY29tMIIBIjAN\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtmUC3p/7EpypE+Dz+gNClvpiw/my\nFFNwBlVRvCZ7ZEGXXLIg2XQnJdjAcc8dGgJcivDvVEWLSzXZyi4tBL02deIXDYRm\nHrUrSCLB449EjjAeCrqhC904BIUDbVV6+caK5y0gTsSidetCAy1G1BMt6INDdjBR\nfsYntvhLGMLL5uyZ38dinGfwmBz+nisgQ1TcQNaec5gsLb0KuO/RB3+bciOMDghh\njkSyfacEaiYfdIB3NnWu9uXK/C68KN6440q9gujv9oQ6fFrGhiu1bTlT8sTmjV5o\n0SuzI9Q8xpxJ4Iiq4L+4JUWDtNRLLx+V2RSTrNgGPLfqLEcX02cuPdZIuwIDAQAB\no4IDHjCCAxowDgYDVR0PAQH/BAQDAgWgMIGgBggrBgEFBQcBAQSBkzCBkDBNBggr\nBgEFBQcwAoZBaHR0cDovL3NlY3VyZS5nbG9iYWxzaWduLmNvbS9jYWNlcnQvZ3Nv\ncmdhbml6YXRpb252YWxzaGEyZzJyMS5jcnQwPwYIKwYBBQUHMAGGM2h0dHA6Ly9v\nY3NwMi5nbG9iYWxzaWduLmNvbS9nc29yZ2FuaXphdGlvbnZhbHNoYTJnMjBWBgNV\nHSAETzBNMEEGCSsGAQQBoDIBFDA0MDIGCCsGAQUFBwIBFiZodHRwczovL3d3dy5n\nbG9iYWxzaWduLmNvbS9yZXBvc2l0b3J5LzAIBgZngQwBAgIwCQYDVR0TBAIwADAh\nBgNVHREEGjAYggsqLnVkZW15LmNvbYIJdWRlbXkuY29tMB0GA1UdJQQWMBQGCCsG\nAQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQU0bmIONs2Ecw4IH2t0IcuXW5bJK0w\nHwYDVR0jBBgwFoAUlt5h8b0cFilTHMDMfTuDAEDmGnwwggF+BgorBgEEAdZ5AgQC\nBIIBbgSCAWoBaAB3AFWB1MIWkDYBSuoLm1c8U/DA5Dh4cCUIFy+jqh0HE9MMAAAB\nao0xIY8AAAQDAEgwRgIhAPkfrrD0CxdAwswgWf6jhT7dd+ta1U9o52bH6GMrpFpP\nAiEA31wxSDDCZIGAg+x5qbOOwkCuw3RzfxN+VpLEAkfd9S8AdgCkuQmQtBhYFIe7\nE6LMZ3AKPDWYBPkb37jjd80OyA3cEAAAAWqNMSF2AAAEAwBHMEUCIFE5yJaIfeyI\n+RTY22NqhOevkAANFV58muomn6BRhAzPAiEA0XqEVHeupxab7vYFjWQvOIKtib1e\nQW7eUqmmJ/0ltywAdQBvU3asMfAxGdiZAKRRFf93FRwR2QLBACkGjbIImjfZEwAA\nAWqNMSHHAAAEAwBGMEQCIEdNV0ULjMdB9ssjYaX2uRdgNE4x3TWCFVgUAj3XA3qT\nAiA6VXlQ8gRTbe4GCc/EqF6mBEASHi34zDTTNws440gMzjANBgkqhkiG9w0BAQsF\nAAOCAQEAGbwsHU9PlTagyC/eH7G5XhuzcDbkTLUg7e2nvy6a1dGtDXhmV19NoGc/\nuwph/E1KTsr9gtp1IDl1U+tgEt62YnLlvJM1lq4mQ5UVUUW+AMI8+PNs/4Mh1N5s\nLXPWPoqpp6vvgo9+zZ5Si6rYwCdILqh5lhv4FRwKQOYInbfIazgOoIwEIKa3DBYK\nXMkE/+JgkXzAeZXyZroZMEzmEHA0NcGy6Yw9CY3/xsJsdlS2dNXJgIMUMZ90Fvkg\nD1QapIZqWwlGLdV+ErUyJpycw6umzd0059f5vLblSOizPLJT5086GX1r/LaTE5ED\nt0SjiH8QPqb5z5KRspazvD9ijH3+BA==\n-----END CERTIFICATE-----\n"]}'
```
### getCert
```
peer chaincode invoke -C mychannel -n certChain -c '{"Args":["getCert","www.shenshimen.com"]}'
```
### modifyCert
```
peer chaincode invoke -C mychannel -n certChain -c '{"Args":["modifyCert","www.shenshimen.com","10597bcc2931a7ed3a29f6709750f367fd68b928fa8c2d2205d979f6dd1d6a49","SHA256","true","-----BEGIN CERTIFICATE-----\nMIIGcjCCBVqgAwIBAgIMBgsWXop3UaR6Fo1hMA0GCSqGSIb3DQEBCwUAMGYxCzAJ\nBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9iYWxTaWduIG52LXNhMTwwOgYDVQQDEzNH\nbG9iYWxTaWduIE9yZ2FuaXphdGlvbiBWYWxpZGF0aW9uIENBIC0gU0hBMjU2IC0g\nRzIwHhcNMTkwNTA2MTI0OTI2WhcNMjAwOTIzMjI1OTM4WjBmMQswCQYDVQQGEwJV\nUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEU\nMBIGA1UEChMLVWRlbXksIEluYy4xFDASBgNVBAMMCyoudWRlbXkuY29tMIIBIjAN\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtmUC3p/7EpypE+Dz+gNClvpiw/my\nFFNwBlVRvCZ7ZEGXXLIg2XQnJdjAcc8dGgJcivDvVEWLSzXZyi4tBL02deIXDYRm\nHrUrSCLB449EjjAeCrqhC904BIUDbVV6+caK5y0gTsSidetCAy1G1BMt6INDdjBR\nfsYntvhLGMLL5uyZ38dinGfwmBz+nisgQ1TcQNaec5gsLb0KuO/RB3+bciOMDghh\njkSyfacEaiYfdIB3NnWu9uXK/C68KN6440q9gujv9oQ6fFrGhiu1bTlT8sTmjV5o\n0SuzI9Q8xpxJ4Iiq4L+4JUWDtNRLLx+V2RSTrNgGPLfqLEcX02cuPdZIuwIDAQAB\no4IDHjCCAxowDgYDVR0PAQH/BAQDAgWgMIGgBggrBgEFBQcBAQSBkzCBkDBNBggr\nBgEFBQcwAoZBaHR0cDovL3NlY3VyZS5nbG9iYWxzaWduLmNvbS9jYWNlcnQvZ3Nv\ncmdhbml6YXRpb252YWxzaGEyZzJyMS5jcnQwPwYIKwYBBQUHMAGGM2h0dHA6Ly9v\nY3NwMi5nbG9iYWxzaWduLmNvbS9nc29yZ2FuaXphdGlvbnZhbHNoYTJnMjBWBgNV\nHSAETzBNMEEGCSsGAQQBoDIBFDA0MDIGCCsGAQUFBwIBFiZodHRwczovL3d3dy5n\nbG9iYWxzaWduLmNvbS9yZXBvc2l0b3J5LzAIBgZngQwBAgIwCQYDVR0TBAIwADAh\nBgNVHREEGjAYggsqLnVkZW15LmNvbYIJdWRlbXkuY29tMB0GA1UdJQQWMBQGCCsG\nAQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQU0bmIONs2Ecw4IH2t0IcuXW5bJK0w\nHwYDVR0jBBgwFoAUlt5h8b0cFilTHMDMfTuDAEDmGnwwggF+BgorBgEEAdZ5AgQC\nBIIBbgSCAWoBaAB3AFWB1MIWkDYBSuoLm1c8U/DA5Dh4cCUIFy+jqh0HE9MMAAAB\nao0xIY8AAAQDAEgwRgIhAPkfrrD0CxdAwswgWf6jhT7dd+ta1U9o52bH6GMrpFpP\nAiEA31wxSDDCZIGAg+x5qbOOwkCuw3RzfxN+VpLEAkfd9S8AdgCkuQmQtBhYFIe7\nE6LMZ3AKPDWYBPkb37jjd80OyA3cEAAAAWqNMSF2AAAEAwBHMEUCIFE5yJaIfeyI\n+RTY22NqhOevkAANFV58muomn6BRhAzPAiEA0XqEVHeupxab7vYFjWQvOIKtib1e\nQW7eUqmmJ/0ltywAdQBvU3asMfAxGdiZAKRRFf93FRwR2QLBACkGjbIImjfZEwAA\nAWqNMSHHAAAEAwBGMEQCIEdNV0ULjMdB9ssjYaX2uRdgNE4x3TWCFVgUAj3XA3qT\nAiA6VXlQ8gRTbe4GCc/EqF6mBEASHi34zDTTNws440gMzjANBgkqhkiG9w0BAQsF\nAAOCAQEAGbwsHU9PlTagyC/eH7G5XhuzcDbkTLUg7e2nvy6a1dGtDXhmV19NoGc/\nuwph/E1KTsr9gtp1IDl1U+tgEt62YnLlvJM1lq4mQ5UVUUW+AMI8+PNs/4Mh1N5s\nLXPWPoqpp6vvgo9+zZ5Si6rYwCdILqh5lhv4FRwKQOYInbfIazgOoIwEIKa3DBYK\nXMkE/+JgkXzAeZXyZroZMEzmEHA0NcGy6Yw9CY3/xsJsdlS2dNXJgIMUMZ90Fvkg\nD1QapIZqWwlGLdV+ErUyJpycw6umzd0059f5vLblSOizPLJT5086GX1r/LaTE5ED\nt0SjiH8QPqb5z5KRspazvD9ijH3+BA==\n-----END CERTIFICATE-----\n"]}'
```
### deleteCert
```
peer chaincode invoke -C mychannel -n certChain -c '{"Args":["deleteCert","www.shenshimen.com"]}'
```

### getHistoryAboutCert
```
peer chaincode invoke -C mychannel -n certChain -c '{"Args":["getHistoryByKey","www.shenshimen.com"]}'
```


