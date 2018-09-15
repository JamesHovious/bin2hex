# bin2hex

Bin2Hex is a command line tool that will parse a binary file and write out a a hex file in an escaped format (e.g. "\x90\41").

```
For example .\Bin2Hex.exe -i MyBytes.bin -o MyHex.hex

Usage: 
  -i, --InFile string    Path to the input file.
  -o, --OutFile string   Path to the output file.
  -l, --linebreak        Multiline output.
  ```