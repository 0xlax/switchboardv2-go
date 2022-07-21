package switchboardv2

import "github.com/gagliardetto/solana-go"

type Env struct {
	Program solana.PublicKey // Program ID
	Mapping solana.PublicKey // Root mapping key
}

var Devnet = Env{
	Program: solana.MustPublicKeyFromBase58(""),
	Mapping: solana.MustPublicKeyFromBase58(""),
}

var Testnet = Env{
	Program: solana.MustPublicKeyFromBase58(""),
	Mapping: solana.MustPublicKeyFromBase58(" "),
}

var Mainnet = Env{
	Program: solana.MustPublicKeyFromBase58(" "),
	Mapping: solana.MustPublicKeyFromBase58(" "),
}
