package storage

import (
	"bytes"
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
)

type Transfer struct {
	Timestamp uint64
	From      common.Address
	To        common.Address
	Value     *big.Int
}

func (s *Storage) AddTransfer(transfer *Transfer) error {
	if !bytes.Equal(transfer.From[:], transfer.To[:]) {
		if err := s.addSingleTransfer(&transfer.From, transfer); err != nil {
			return err
		}
		if err := s.addSingleTransfer(&transfer.To, transfer); err != nil {
			return err
		}
	} else {
		if err := s.addSingleTransfer(&transfer.To, transfer); err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) GetTransfer(address *common.Address, count uint64) (*Transfer, error) {

	key := []byte(prefixTransfer)
	key = append(key, address.Bytes()...)
	key = append(key, uint642bytes(count)...)

	value, err := s.db.Get(key, nil)
	if err != nil {
		return nil, err
	}

	var entry Transfer
	err = rlp.DecodeBytes(value, &entry)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (s *Storage) GetTransferCount(address *common.Address) (uint64, error) {
	key := []byte(prefixTransferCount)
	key = append(key, address.Bytes()...)

	value, err := s.db.Get(key, nil)
	if err == leveldb.ErrNotFound {
		value = make([]byte, 8)
	} else if err != nil {
		return 0, err
	}

	return bytes2uint64(value), nil
}

func (s *Storage) addSingleTransfer(address *common.Address, transfer *Transfer) error {
	count, err := s.GetTransferCount(address)
	if err != nil {
		return nil
	}

	key := []byte(prefixTransfer)
	key = append(key, address.Bytes()...)
	key = append(key, uint642bytes(count)...)

	var value []byte
	if value, err = rlp.EncodeToBytes(transfer); err != nil {
		return err
	}

	err = s.db.Put(key, value, nil)
	if err != nil {
		return err
	}

	return s.putTransferCount(address, count+1)
}

func (s *Storage) putTransferCount(address *common.Address, count uint64) error {
	key := []byte(prefixTransferCount)
	key = append(key, address.Bytes()...)
	return s.db.Put(key, uint642bytes(count), nil)
}

func bytes2uint64(value []byte) uint64 {
	return binary.BigEndian.Uint64(value)
}

func uint642bytes(value uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, value)
	return bytes
}
