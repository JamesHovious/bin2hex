package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	buffer.WriteRune('\\')
	buffer.WriteRune('x')
	var n1 = n - 1
	var l1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n1 && i != l1 {
			buffer.WriteRune('\\')
			buffer.WriteRune('x')
		}
	}
	return buffer.String()
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 2 {
		fmt.Println("[!] Should only have two arguments!")
		os.Exit(0)
	}
	bytes, err := ioutil.ReadFile(argsWithoutProg[0])
	if err != nil {
		fmt.Println("[!] Error reading file")
		os.Exit(0)
	}
	binStr := hex.EncodeToString(bytes)
	hexBytes := []byte(insertNth(binStr, 2))

	err = ioutil.WriteFile(argsWithoutProg[1], hexBytes, 0644)
	if err != nil {
		fmt.Println("[+] Successfully wrote out file")
	}

}
