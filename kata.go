package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"runtime"
)

const maximum = 99999999

func Crack(hash string) string {
	for i := 0; i <= maximum; i += 1 {
		paddedString := fmt.Sprintf("%05d", i)
		currentHash := md5.Sum([]byte(paddedString))
		currentHashString := hex.EncodeToString(currentHash[:])
		if currentHashString == hash {
			return paddedString
		}
	}
	return ""
}

func md5Hash(s string) string {
	currentHash := md5.Sum([]byte(s))
	return hex.EncodeToString(currentHash[:])
}

func main() {
	s := md5Hash("99999999")
	fmt.Println(s)
	fmt.Println(runtime.NumCPU())
}
