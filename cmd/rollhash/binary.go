package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func writeBinaryFile(fileName string, data map[int][]byte) {
	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	for key, value := range data {
		// Write the key to the output file
		if err := binary.Write(outputFile, binary.LittleEndian, int32(key)); err != nil {
			fmt.Println("Error writing key to output file:", err)
			return
		}

		// Write the length of the value slice followed by the value slice itself
		valueLength := int32(len(value))
		if err := binary.Write(outputFile, binary.LittleEndian, valueLength); err != nil {
			fmt.Println("Error writing value length to output file:", err)
			return
		}
		if _, err := outputFile.Write(value); err != nil {
			fmt.Println("Error writing value to output file:", err)
			return
		}
	}
}

func readBinaryFile(fileName string) map[int][]byte {
	data := make(map[int][]byte)

	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return data
	}
	defer inputFile.Close()

	var key int32
	for {
		// Read the key from the input file
		if err := binary.Read(inputFile, binary.LittleEndian, &key); err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading key from input file:", err)
			return data
		}

		// Read the length of the value slice from the input file
		var valueLength int32
		if err := binary.Read(inputFile, binary.LittleEndian, &valueLength); err != nil {
			fmt.Println("Error reading value length from input file:", err)
			return data
		}

		// Read the value slice from the input file
		value := make([]byte, valueLength)
		if _, err := inputFile.Read(value); err != nil {
			fmt.Println("Error reading value from input file:", err)
			return data
		}

		data[int(key)] = value
	}

	return data
}
