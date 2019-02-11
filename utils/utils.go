package utils

import (
	"bytes"
	"encoding/binary"
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

func Uint64ToEthBytes(u uint64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, u)
	if err != nil {
		panic(err)
	}
	var r [32]byte
	copy(r[32-len(buff.Bytes()):], buff.Bytes())
	return r[:]
}
