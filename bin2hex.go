package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var inFilePtr, outFilePtr *string
var hexBytes []byte

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

func parseArgs() {
	inFilePtr = flag.String("InFile", "", " Where is the input file.")
	outFilePtr = flag.String("Outfile", "", " Where do you want to save the file.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\r\n"+
			"Bin2Hex is a command line tool that will parse a binary file and write out a a hex file in an escaped format (e.g. \"\\x90\\41\").\r\n\r\n"+
			"For example .\\Bin2Hex.exe --InFile MyBytes.bin --OutFile MyHex.hex\r\n\r\n"+
			"Usage: \r\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()
}

func readInputfile() {
	// Get the args and work with them.
	bytes, err := ioutil.ReadFile(*inFilePtr)
	if err != nil {
		fmt.Println("[!] Error reading file")
		fmt.Println(err)
		os.Exit(0)
	}
	binStr := hex.EncodeToString(bytes)
	hexBytes = []byte(insertNth(binStr, 2))
}

func writeOutputFile() {
	err := ioutil.WriteFile(*outFilePtr, hexBytes, 0644)
	if err != nil {
		fmt.Println("[!] Error writing file")
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("[+] Successfully wrote out file")
}

func main() {
	parseArgs()
	readInputfile()
	writeOutputFile()
}
