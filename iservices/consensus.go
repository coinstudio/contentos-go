package iservices

import (
	"github.com/coschain/contentos-go/common"
)

var ConsensusServerName = "consensus"

type IConsensus interface {
	// NOTE: producers should be maintained by the specific Consensus algorithm
	// CurrentProducer returns current producer
	CurrentProducer() string
	// ActiveProducers returns a list of accounts that actively produce blocks
	ActiveProducers() []string

	// SetBootstrap determines if the current node starts a new block chain
	SetBootstrap(b bool)

	// PushTransaction accepts the trx
	PushTransaction(trx common.ISignedTransaction, wait bool, broadcast bool) common.ITransactionReceiptWithInfo

	// PushBlock adds b to the block fork DB, called if ValidateBlock returns true
	PushBlock(b common.ISignedBlock)

	// Push sends a user defined msg to consensus
	Push(msg interface{})

	// GetLastBFTCommit get the last irreversible block info. @evidence
	// is the information that can prove the id is indeed the last irreversible one.
	// e.g. if user uses bft to achieve fast ack, @evidence can simply be the collection
	// of the vote message
	GetLastBFTCommit() (evidence interface{})

	// GetHeadBlockId returns the block id of the head block
	GetHeadBlockId() common.BlockID

	// GetIDs returns a list of block ids which remote peer may not have
	GetIDs(start, end common.BlockID) ([]common.BlockID, error)

	// FetchBlock returns the block whose id is the given param
	FetchBlock(id common.BlockID) (common.ISignedBlock, error)

	// HasBlock query the local blockchain whether it has the given block id
	HasBlock(id common.BlockID) bool

	FetchBlocksSince(id common.BlockID) ([]common.ISignedBlock, error)
}
