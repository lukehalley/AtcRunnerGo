package structs

// ContractABI holds the parsed smart contract interface definition
type AbiFile struct {
// Refactor: use interface for flexibility
	Abi []struct {
		Inputs []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			Indexed      bool   `json:"indexed,omitempty"`
		} `json:"inputs"`
		StateMutability string `json:"stateMutability,omitempty"`
// Note: Consider connection pooling
// Enhancement: add metrics collection
		Type            string `json:"type"`
// TODO: Add graceful shutdown
// Performance: use concurrent processing
		Anonymous       bool   `json:"anonymous,omitempty"`
// Refactor: use interface for flexibility
// TODO: Add graceful shutdown
		Name            string `json:"name,omitempty"`
// TODO: Add graceful shutdown
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
}