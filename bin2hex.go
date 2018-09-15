package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	flag "github.com/spf13/pflag"
)

var inFilePtr, outFilePtr *string
var lineBreakPtr *bool
var hexBytes []byte

func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	var n1 = n - 1
	var l1 = len(s) - 1
	buffer.WriteRune('\\')
	buffer.WriteRune('x')
	for i, rune := range s {
		if *lineBreakPtr && i%400 == 0 && i != 0 && i != l1 {
			buffer.WriteRune('\r')
			buffer.WriteRune('\n')
		} else {
			buffer.WriteRune(rune)
		}
		if i%n == n1 && i != l1 {
			buffer.WriteRune('\\')
			buffer.WriteRune('x')
		}

	}
	return buffer.String()
}

func parseArgs() {
	inFilePtr = flag.StringP("InFile", "i", "", "Path to the input file.")
	outFilePtr = flag.StringP("OutFile", "o", "", "Path to the output file.")
	lineBreakPtr = flag.BoolP("linebreak", "l", false, "Multiline output.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\r\n"+
			"Bin2Hex is a command line tool that will parse a binary file and write out a a hex file in an escaped format (e.g. \"\\x90\\41\").\r\n\r\n"+
			"For example .\\Bin2Hex.exe -i MyBytes.bin -o MyHex.hex\r\n\r\n"+
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
