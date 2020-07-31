package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	//Read the certificate from *.pem
	data, err := ioutil.ReadFile("./certificates/udemy-com.pem")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//Compute the hashcode using SHA256
	message := []byte(string(data))
	hashcode := GetSHA256HashCode(message)
	fmt.Println("certificate", string(data))
	fmt.Println("SHA256 hash:", hashcode)
}

//SHA256生成哈希值
func GetSHA256HashCode(message []byte) string {
	//计算哈希值，返回一个长度为32的数组
	bytes := sha256.Sum256(message)
	//将数组转换成切片，转换成16进制，返回字符串
	hashcode := hex.EncodeToString(bytes[:])
	return hashcode
}
