# Blockchain for certificate management
## Installation & Initialization
To install:
```
peer chaincode install -n certChain -p github.com/hyperledger/fabric/singlepeer/chaincode/go/CertChain/ -v 1.0
```
To instantiate:
```
peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n certChain -v 1.0 -c '{"Args":["init","CertChain","a blockchain for certificate management"]}' -P "AND ('Org1MSP.peer')"
```
## Test
### addCert
```
peer chaincode invoke -C mychannel -n certChain -c '{"Args":["addCert","www.zhaizhonghao.com","40fceca6612099c3eed324e1cc701ac59f01067b45ccf484ffd4d6e9df16232c","SHA256","false","-----BEGIN CERTIFICATE-----\nMIIGcjCCBVqgAwIBAgIMBgsWXop3UaR6Fo1hMA0GCSqGSIb3DQEBCwUAMGYxCzAJ\nBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9iYWxTaWduIG52LXNhMTwwOgYDVQQDEzNH\nbG9iYWxTaWduIE9yZ2FuaXphdGlvbiBWYWxpZGF0aW9uIENBIC0gU0hBMjU2IC0g\nRzIwHhcNMTkwNTA2MTI0OTI2WhcNMjAwOTIzMjI1OTM4WjBmMQswCQYDVQQGEwJV\nUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEU\nMBIGA1UEChMLVWRlbXksIEluYy4xFDASBgNVBAMMCyoudWRlbXkuY29tMIIBIjAN\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtmUC3p/7EpypE+Dz+gNClvpiw/my\nFFNwBlVRvCZ7ZEGXXLIg2XQnJdjAcc8dGgJcivDvVEWLSzXZyi4tBL02deIXDYRm\nHrUrSCLB449EjjAeCrqhC904BIUDbVV6+caK5y0gTsSidetCAy1G1BMt6INDdjBR\nfsYntvhLGMLL5uyZ38dinGfwmBz+nisgQ1TcQNaec5gsLb0KuO/RB3+bciOMDghh\njkSyfacEaiYfdIB3NnWu9uXK/C68KN6440q9gujv9oQ6fFrGhiu1bTlT8sTmjV5o\n0SuzI9Q8xpxJ4Iiq4L+4JUWDtNRLLx+V2RSTrNgGPLfqLEcX02cuPdZIuwIDAQAB\no4IDHjCCAxowDgYDVR0PAQH/BAQDAgWgMIGgBggrBgEFBQcBAQSBkzCBkDBNBggr\nBgEFBQcwAoZBaHR0cDovL3NlY3VyZS5nbG9iYWxzaWduLmNvbS9jYWNlcnQvZ3Nv\ncmdhbml6YXRpb252YWxzaGEyZzJyMS5jcnQwPwYIKwYBBQUHMAGGM2h0dHA6Ly9v\nY3NwMi5nbG9iYWxzaWduLmNvbS9nc29yZ2FuaXphdGlvbnZhbHNoYTJnMjBWBgNV\nHSAETzBNMEEGCSsGAQQBoDIBFDA0MDIGCCsGAQUFBwIBFiZodHRwczovL3d3dy5n\nbG9iYWxzaWduLmNvbS9yZXBvc2l0b3J5LzAIBgZngQwBAgIwCQYDVR0TBAIwADAh\nBgNVHREEGjAYggsqLnVkZW15LmNvbYIJdWRlbXkuY29tMB0GA1UdJQQWMBQGCCsG\nAQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQU0bmIONs2Ecw4IH2t0IcuXW5bJK0w\nHwYDVR0jBBgwFoAUlt5h8b0cFilTHMDMfTuDAEDmGnwwggF+BgorBgEEAdZ5AgQC\nBIIBbgSCAWoBaAB3AFWB1MIWkDYBSuoLm1c8U/DA5Dh4cCUIFy+jqh0HE9MMAAAB\nao0xIY8AAAQDAEgwRgIhAPkfrrD0CxdAwswgWf6jhT7dd+ta1U9o52bH6GMrpFpP\nAiEA31wxSDDCZIGAg+x5qbOOwkCuw3RzfxN+VpLEAkfd9S8AdgCkuQmQtBhYFIe7\nE6LMZ3AKPDWYBPkb37jjd80OyA3cEAAAAWqNMSF2AAAEAwBHMEUCIFE5yJaIfeyI\n+RTY22NqhOevkAANFV58muomn6BRhAzPAiEA0XqEVHeupxab7vYFjWQvOIKtib1e\nQW7eUqmmJ/0ltywAdQBvU3asMfAxGdiZAKRRFf93FRwR2QLBACkGjbIImjfZEwAA\nAWqNMSHHAAAEAwBGMEQCIEdNV0ULjMdB9ssjYaX2uRdgNE4x3TWCFVgUAj3XA3qT\nAiA6VXlQ8gRTbe4GCc/EqF6mBEASHi34zDTTNws440gMzjANBgkqhkiG9w0BAQsF\nAAOCAQEAGbwsHU9PlTagyC/eH7G5XhuzcDbkTLUg7e2nvy6a1dGtDXhmV19NoGc/\nuwph/E1KTsr9gtp1IDl1U+tgEt62YnLlvJM1lq4mQ5UVUUW+AMI8+PNs/4Mh1N5s\nLXPWPoqpp6vvgo9+zZ5Si6rYwCdILqh5lhv4FRwKQOYInbfIazgOoIwEIKa3DBYK\nXMkE/+JgkXzAeZXyZroZMEzmEHA0NcGy6Yw9CY3/xsJsdlS2dNXJgIMUMZ90Fvkg\nD1QapIZqWwlGLdV+ErUyJpycw6umzd0059f5vLblSOizPLJT5086GX1r/LaTE5ED\nt0SjiH8QPqb5z5KRspazvD9ijH3+BA==\n-----END CERTIFICATE-----\n"]}'
```
### getCert
```
peer chaincode invoke -C mychannel -n certChain -c '{"Args":["getCert","www.zhaizhonghao.com"]}'
```
### modifyCert
```
peer chaincode install -n certChain -p github.com/hyperledger/fabric/singlepeer/chaincode/go/CertChain/ -v 1.0
```
### deleteCert
```
peer chaincode install -n certChain -p github.com/hyperledger/fabric/singlepeer/chaincode/go/CertChain/ -v 1.0
```

