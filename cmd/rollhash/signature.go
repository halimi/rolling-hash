package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/halimi/rolling-hash/signature"
)

const (
	minChunkSize  = 1
	maxChunkSize  = 1024
	chunkSizeHash = "chunksize"
)

func signatureCmd(args []string) {
	argNum := len(args)
	if argNum < 1 || argNum > 2 {
		fmt.Println("Invalid argument numbers.")
		return
	}

	inputFileName := args[0]
	outputFileName := ""

	if len(args) == 2 {
		outputFileName = args[1]
	}

	inpFile, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	chunkSize := len(inpFile) / 2
	if chunkSize < minChunkSize {
		chunkSize = minChunkSize
	}
	if chunkSize > maxChunkSize {
		chunkSize = maxChunkSize
	}

	sig := signature.NewSignature(chunkSize)
	signatures := sig.Signatures(inpFile)

	// Add the chunk size to the signatures
	signatures[chunkSizeHash] = []int{chunkSize}

	if outputFileName == "" {
		fmt.Printf("%v\n", signatures)
	} else {
		outFile, err := os.Create(outputFileName)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outFile.Close()

		// Encode data to JSON and write to outputFile
		encoder := json.NewEncoder(outFile)
		if err := encoder.Encode(signatures); err != nil {
			fmt.Println("Error encoding data to JSON:", err)
			return
		}
	}
}
