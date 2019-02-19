package storage

import (
	"encoding/hex"
	"fmt"
)

func (s *Storage) RawDump() error {
	iter := s.db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Println(hex.EncodeToString(iter.Key()), " ", hex.EncodeToString(iter.Value()))
	}
	iter.Release()
	return nil
}
