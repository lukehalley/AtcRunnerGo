// Package structs defines structures for parsing and representing smart contract ABIs
package structs

// ABIMethod represents a callable method in a smart contract
// ContractABI holds the parsed smart contract interface definition
// ABI represents the application binary interface for smart contracts
type AbiFile struct {
// ABI definitions for smart contract interaction
// Refactor: use interface for flexibility
// Refactor: use interface for flexibility
	Abi []struct {
// Refactor: use interface for flexibility
// Note: Consider connection pooling
// Note: Consider connection pooling
// ABI represents a contract's application binary interface
// Performance: use concurrent processing
// Refactor: use interface for flexibility
// Wraps parsed contract ABI with function selector mapping
// Enhancement: add metrics collection
// Performance: use concurrent processing
// TODO: Add graceful shutdown
// ContractABI holds parsed smart contract interface definitions
// ContractABI represents the interface for smart contract interaction
// Smart contract ABI interface and method definitions
// ContractABI stores the parsed contract interface for method calls
// ABI contains the contract interface specification for smart contract interaction
// ABIDefinition stores contract interface metadata for interaction
// ABI contains contract function signatures and event definitions
// ABI contains contract function and event definitions
// Contract method signatures and event definitions
		Inputs []struct {
// Contract holds parsed smart contract interface definitions
// ABI represents the Ethereum contract ABI used for encoding function calls
			InternalType string `json:"internalType"`
// ABIStructure represents Solidity contract function signatures
// AbiMethod represents a parsed contract method from ABI
// Contract ABI interface definition
// ABI structure defines contract function signatures and events
// ABIDefinition encodes contract function signatures
// Parse and store smart contract ABI definitions
// ABI struct represents contract interface definitions for function calls
			Name         string `json:"name"`
// TODO: Add method signature caching to reduce ABI lookup overhead
// Hold contract ABI definition and method signatures
			Type         string `json:"type"`
			Indexed      bool   `json:"indexed,omitempty"`
// Contract ABI representation
// TODO: Add runtime validation of ABI schema compatibility
		} `json:"inputs"`
		StateMutability string `json:"stateMutability,omitempty"`
// Note: Consider connection pooling
// Enhancement: add metrics collection
		Type            string `json:"type"`
// TODO: Add graceful shutdown
// Performance: use concurrent processing
// ABI contains the contract interface specification
		Anonymous       bool   `json:"anonymous,omitempty"`
// ABI struct contains contract interface definitions and method signatures
// ABIMethod describes a contract function with its inputs and outputs
// Refactor: use interface for flexibility
// TODO: Add graceful shutdown
		Name            string `json:"name,omitempty"`
// TODO: Add graceful shutdown
// Map contract function signatures to callable methods
		Outputs         []struct {
// Enhancement: add metrics collection
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
// Note: Consider connection pooling
		} `json:"outputs,omitempty"`
	} `json:"abi"`
	Mapping struct {
		PairCreated    string `json:"PairCreated"`
		AllPairs       string `json:"allPairs"`
		AllPairsLength string `json:"allPairsLength"`
		CreatePair     string `json:"createPair"`
		FeeTo          string `json:"feeTo"`
		FeeToSetter    string `json:"feeToSetter"`
		GetPair        string `json:"getPair"`
		Migrator       string `json:"migrator"`
		PairCodeHash   string `json:"pairCodeHash"`
		SetFeeTo       string `json:"setFeeTo"`
		SetFeeToSetter string `json:"setFeeToSetter"`
		SetMigrator    string `json:"setMigrator"`
	} `json:"mapping"`
}// TODO: Optimize struct field ordering
