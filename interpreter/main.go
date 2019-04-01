package main

import (
	"github.com/kostyaBro/SICP/interpreter/LineReader"
	"os"
)

var (
	help = "LISP interpreter\n" +
		"--help, -h    show documentation\n" +
		"exit          close interpreter"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help":
		case "-h":
			println(help)
			break
		case "--compile":
		case "-cl":
			// todo compile to binary
		case "--run":
		case "-r":
			// todo runner
		case "--interpretToGolang":
		case "-itg":
			// todo interpret
		default:
			// todo command %s not found
		}
	} else {
		// todo LineReader
		heep := new(LineReader.Heep)
		lineReader := LineReader.New(heep)
		LineReader.NewProcessor(lineReader).Process()
	}
}
