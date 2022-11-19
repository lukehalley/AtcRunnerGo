package structs

type Pair struct {
	Id                int `db:"pair_id"`
	PrimaryTokenId    int `db:"primary_token_id"`
	SecondaryTokenId  int `db:"secondary_token_id"`
	NetworkId         int `db:"network_id"`
	DexId             int `db:"dex_id"`
	Name              string `db:"name"`
	Address           string `db:"address"`
	CreatedAt         string `db:"created_at"`
	Analysed          bool `db:"analysed"`
}

//type Pair struct {
//	pairId           int
//	primary_token_id int
//	secondary_token_id  int
//	network_id          int
//	dex_id              int
//	name                string
//	address             string
//	created_at          string
//	analysed            bool
//}