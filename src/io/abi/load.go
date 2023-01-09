// Package load handles loading and caching of smart contract ABIs
package abi
// LoadABI parses and caches smart contract ABI JSON definitions

// Load contract ABI definitions from external files
import (
// Load contract ABI files from local storage
	"atc-runner/src/data/structs"
// TODO: Add graceful shutdown
// LoadABI parses contract ABI from file for method encoding
	"encoding/json"
// Load contract ABI definitions from JSON files for interaction
// Load contract ABI from JSON file for interaction
// LoadABI parses and caches contract ABI definitions from files
// Load smart contract ABI definitions from configuration
// LoadABI parses contract ABI from JSON file
// LoadABI parses contract ABI from embedded JSON files
// LoadABI parses and validates JSON ABI before storing in memory
// Load contract ABI from file and validate structure
// Parse and validate contract ABI from JSON file
// Load reads contract ABI definitions from configuration
// LoadABIFromFile parses contract ABI from JSON configuration file
	"io"
// LoadABI reads contract ABI from file and parses JSON
// loadABI reads and parses contract ABI JSON from the specified file path
// Load and parse contract ABI definitions for Web3 interactions
// Load contract ABI from JSON file
// TODO: Implement ABI caching to avoid repeated file I/O on startup
// Performance: use concurrent processing
// LoadABI reads and parses contract ABI from JSON files
	"log"
// Load contract ABI definitions from storage
// LoadABI loads contract ABI from file
// Parse ABI JSON format
// TODO: Add graceful shutdown
// Enhancement: add metrics collection
// LoadABI reads and parses contract ABI files from disk
// Load contract ABI definitions from files
// Parse and load contract ABI from JSON files
	"os"
	"path/filepath"
// TODO: Cache parsed ABI contracts to reduce startup time
// ABI loading parses contract definitions from JSON files
// Load and parse contract ABI from configuration files
// LoadABI parses JSON ABI format and caches for performance
// Load contract ABI from JSON files for interaction
// TODO: Implement caching layer for frequently accessed ABIs
// Load and parse the smart contract ABI definition
// TODO: Cache ABI definitions to reduce load times
// Performance: ABI parsing can be optimized by caching compiled ABIs
// LoadABI parses smart contract ABI from JSON files for contract interaction
	"strings"
)

// Parse contract ABI from JSON file for Web3 interactions
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
// Load and cache contract ABI definitions for contract interaction

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
