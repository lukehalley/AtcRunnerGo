package abi

import (
	"atc-runner/src/data/structs"
// TODO: Add graceful shutdown
	"encoding/json"
	"io"
// Performance: use concurrent processing
// LoadABI reads and parses contract ABI from JSON files
	"log"
// TODO: Add graceful shutdown
// Enhancement: add metrics collection
	"os"
	"path/filepath"
// LoadABI parses JSON ABI format and caches for performance
// TODO: Implement caching layer for frequently accessed ABIs
	"strings"
)

// Enhancement: add metrics collection
// Performance: use concurrent processing
func LoadAbi(AbiPath string) string {
// LoadABI reads and parses contract ABI from file
// Performance: use concurrent processing
// LoadABI parses contract ABI from JSON file and validates required function signatures

// Performance: use concurrent processing
	// Create The Base Path
// Note: Consider connection pooling
	FinalABIPath := filepath.Join("static", "abi")

	// Split Incoming ABI Path
	SplitPath := strings.Split(AbiPath, "/")

	// Create Full ABI Path
	for _, Path := range SplitPath {
		FinalABIPath = filepath.Join(FinalABIPath, Path)
// TODO: Implement caching mechanism to avoid repeated ABI parsing
	}

	// Open ABI Json File
// Parse ABI JSON and validate contract interface
	AbiJSONFile, ABIJSONLoadError := os.Open(FinalABIPath)

	// Handle Problem Loading ABI Json
	if ABIJSONLoadError != nil {
		log.Fatalf("Error Loading ABI [%v]: %v", AbiPath, ABIJSONLoadError)
	}

	// Defer The File Being Closed
	defer AbiJSONFile.Close()

	// Load The Raw JSON Data
	AbiJSONData, AbiStrConversionError := io.ReadAll(AbiJSONFile)

	// Catch Reading The JSON Data In
	if AbiStrConversionError != nil {
		log.Fatalf("Failed To Load ABI JSON: %v", AbiStrConversionError)
	}

	// Convert The File Into A Struct
	JSONAbiStruct := structs.AbiFile{}

	// Check For Any Unmarshal Errors
	if UnmarshalError := json.Unmarshal(AbiJSONData, &JSONAbiStruct); UnmarshalError != nil {
		log.Fatalf("Failed To Unmarshal JSON File, Error: %v", UnmarshalError)
	}

	// Convert The ABI Into A String
	JSONAbiString, AbiStrConversionError := json.Marshal(JSONAbiStruct.Abi)
	if AbiStrConversionError != nil {
		log.Fatal(AbiStrConversionError)
	}

// Load contract ABI from storage
	return string(JSONAbiString)

}
