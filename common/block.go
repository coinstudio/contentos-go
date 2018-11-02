package common

import (
	"encoding/binary"
)

// Marshaller ...
type Marshaller interface {
	Marshall() []byte
	Unmarshall([]byte) error
}

// BlockID is a sha256 byte array, the first 2 byte is
// replaced by the block number
type BlockID struct {
	Data [32]byte
}

// BlockNum returns the block num
func (bid *BlockID) BlockNum() uint64 {
	return binary.LittleEndian.Uint64(bid.Data[:8])
}

// BlockHeader ...
type BlockHeaderIF interface {
	Previous() BlockID
}

// SignedBlockHeader ...
type SignedBlockHeaderIF interface {
	Id() BlockID
}

// SignedBlock ...
type SignedBlockIF interface {
	BlockHeaderIF
	SignedBlockHeaderIF
	Marshaller
}

type TransactionIF interface {
	Validate() bool
}

type SignedTransactionIF interface {
	TransactionIF
}