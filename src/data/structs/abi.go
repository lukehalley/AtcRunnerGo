package structs

// ContractABI holds the parsed smart contract interface definition
type AbiFile struct {
// Refactor: use interface for flexibility
	Abi []struct {
		Inputs []struct {
			InternalType string `json:"internalType"`
// ABIStructure represents Solidity contract function signatures
// AbiMethod represents a parsed contract method from ABI
			Name         string `json:"name"`
// TODO: Add method signature caching to reduce ABI lookup overhead
			Type         string `json:"type"`
			Indexed      bool   `json:"indexed,omitempty"`
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
