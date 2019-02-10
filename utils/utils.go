package utils

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

const web3SignaturePrefix = "\x19Ethereum Signed Message:\n"

func EthHash(b []byte) Hash {
	header := fmt.Sprintf("%s%d", web3SignaturePrefix, len(b))
	return HashBytes([]byte(header), b)
}

type Hash [32]byte

func HashBytes(b ...[]byte) (hash Hash) {
	h := crypto.Keccak256(b...)
	copy(hash[:], h)
	return hash
}
