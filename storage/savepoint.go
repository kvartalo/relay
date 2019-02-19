package storage

import (
	"github.com/ethereum/go-ethereum/rlp"
)

type SavePoint struct {
	LastBlock   uint64
	LastTxIndex uint
}

func (s *Storage) LoadSavePoint() (*SavePoint, error) {
	key := []byte(prefixSavePoint)
	value, err := s.db.Get(key, nil)
	if err != nil {
		return nil, err
	}

	var entry SavePoint
	err = rlp.DecodeBytes(value, &entry)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

// SetGlobals in the storage.
func (s *Storage) SetSavePoint(savePoint SavePoint) error {

	var err error

	key := []byte(prefixSavePoint)
	var value []byte

	if value, err = rlp.EncodeToBytes(savePoint); err != nil {
		return err
	}

	return s.db.Put(key, value, nil)
}
