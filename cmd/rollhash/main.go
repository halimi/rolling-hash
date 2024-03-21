package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  rollhash signature [BASIS [SIGNATURE]]")
	fmt.Println("  rollhash delta SIGNATURE [NEWFILE [DELTA]]")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No arguments provided.")
		printHelp()
		os.Exit(1)
	}

	switch args[0] {
	case "signature":
		signatureCmd(args[1:])
	case "delta":
		deltaCmd(args[1:])
	default:
		fmt.Println("Unknown command:", args[0])
		printHelp()
		os.Exit(1)
	}
}
