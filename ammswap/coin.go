package ammswap

type Coin struct {
	Decimals  int
	Name      string
	Symbol    string
	TokenType TokenType
}

type TokenType struct {
	Address    string
	Module     string
	StructName string
}
