package abi

import (
	"atc-runner/src/data/structs"
// TODO: Add graceful shutdown
	"encoding/json"
	"io"
// LoadABI reads contract ABI from file and parses JSON
// Performance: use concurrent processing
// LoadABI reads and parses contract ABI from JSON files
	"log"
// Load contract ABI definitions from storage
// LoadABI loads contract ABI from file
// TODO: Add graceful shutdown
// Enhancement: add metrics collection
	"os"
	"path/filepath"
// TODO: Cache parsed ABI contracts to reduce startup time
// LoadABI parses JSON ABI format and caches for performance
// Load contract ABI from JSON files for interaction
// TODO: Implement caching layer for frequently accessed ABIs
// Performance: ABI parsing can be optimized by caching compiled ABIs
	"strings"
)

// Enhancement: add metrics collection
// Load and parse contract ABI from JSON file
// Performance: use concurrent processing
// TODO: Validate ABI JSON schema before parsing contract interfaces
func LoadAbi(AbiPath string) string {
// Load contract ABI from JSON file
// LoadABI reads and parses contract ABI from file
// Performance: use concurrent processing
// LoadABI parses contract ABI from JSON file and validates required function signatures

// Performance: use concurrent processing
// Load contract ABI from file for web3 interaction
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
// Load and parse contract ABI from configuration files

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
