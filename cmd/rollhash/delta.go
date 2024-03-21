package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/halimi/rolling-hash/delta"
)

func deltaCmd(args []string) {
	argNum := len(args)
	if argNum < 2 || argNum > 3 {
		fmt.Println("Invalid argument numbers.")
		return
	}

	sigFileName := args[0]
	newFileName := args[1]
	deltaFileName := ""

	if len(args) == 3 {
		deltaFileName = args[2]
	}

	sigFile, err := os.Open(sigFileName)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}
	defer sigFile.Close()

	// Decode JSON data from the outputFile into a map[string][]int
	var signatures map[string][]int
	decoder := json.NewDecoder(sigFile)
	if err := decoder.Decode(&signatures); err != nil {
		fmt.Println("Error decoding data from JSON:", err)
		return
	}

	chunkSizeList, ok := signatures[chunkSizeHash]
	if !ok || len(chunkSizeList) == 0 {
		fmt.Println("Chunk size not found.")
		return
	}
	chunkSize := chunkSizeList[0]
	delete(signatures, chunkSizeHash)

	newFile, err := os.ReadFile(newFileName)
	if err != nil {
		fmt.Println("Error reading new file:", err)
		return
	}

	delta := delta.NewDelta(chunkSize)
	deltas, err := delta.Diff(signatures, newFile)
	if err != nil {
		fmt.Println("Error during calculating the diff:", err)
		return
	}

	if deltaFileName == "" {
		fmt.Printf("%v\n", deltas)
	} else {
		writeBinaryFile(deltaFileName, deltas)
	}
}
