package consensus

import (
	"encoding/binary"
	"fmt"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/db/storage"
	"github.com/coschain/gobft/message"
)

/***************************
validators:
	if reach consensus:
		add checkPoint
		if do commit success:
			flush checkPoint
		else if committing a missing block:
			do sync
	recv message.Commit:
		if not committed already:
			pass message.Commit to gobft
non-validators:
	recv message.Commit:
		if not committed already && num within range:
			add checkPoint
			if lib num == (message.Commit).LastCommitted:
				if validate(message.Commit)==true:
					if has the block about to be committed:
						if do commit success:
							flush checkPoint
							commit later blocks if possible
					else:
						do sync
				else:
					remove checkPoint
push block b:
	if b.id == next_checkPoint.id:
		if do commit success:
			flush checkPoint
***************************/

// BFTCheckPoint maintains the bft consensus evidence, the votes collected
// for the same checkpoint in different validators might differ. But all
// nodes including validators should have the same number of checkpoints with
// exact same order.
// all methods have time complexity of O(1)
type BFTCheckPoint struct {
	sabft   *SABFT
	dataDir string
	db      storage.Database

	lastCommitted common.BlockID
	nextCP        common.BlockID
	cache         map[common.BlockID]*message.Commit // lastCommitted-->Commit
}

func NewBFTCheckPoint(dir string, sabft *SABFT) *BFTCheckPoint {
	db, err := storage.NewLevelDatabase(dir)
	if err != nil {
		panic(err)
	}
	lc := sabft.ForkDB.LastCommitted()
	return &BFTCheckPoint{
		sabft:         sabft,
		dataDir:       dir,
		db:            db,
		lastCommitted: lc,
		nextCP:        common.EmptyBlockID,
		cache:         make(map[common.BlockID]*message.Commit),
	}
}

func (cp *BFTCheckPoint) Flush() error {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, cp.nextCP.BlockNum())
	err := cp.db.Put(key, cp.cache[cp.lastCommitted].Bytes())
	if err != nil {
		cp.sabft.log.Fatal(err)
		return err
	}
	delete(cp.cache, cp.lastCommitted)

	cp.lastCommitted = cp.nextCP
	cp.sabft.log.Info("checkpoint flushed at block height ", cp.nextCP.BlockNum())
	cp.nextCP = common.EmptyBlockID
	if v, ok := cp.cache[cp.lastCommitted]; ok {
		cp.nextCP = ConvertToBlockID(v.ProposedData)
	}
	return nil
}

func (cp *BFTCheckPoint) Add(commit *message.Commit) error {
	if err := commit.ValidateBasic(); err != nil {
		cp.sabft.log.Error(err)
		return ErrInvalidCheckPoint
	}
	blockID := ExtractBlockID(commit)
	blockNum := blockID.BlockNum()
	libNum := cp.lastCommitted.BlockNum()
	if blockNum > libNum+constants.MaxUncommittedBlockNum ||
		blockNum <= libNum {
		cp.sabft.log.Error(ErrCheckPointOutOfRange)
		return ErrCheckPointOutOfRange
	}

	prev := ConvertToBlockID(commit.Prev)
	if _, ok := cp.cache[prev]; ok {
		return ErrCheckPointExists
	}
	cp.cache[prev] = commit
	if cp.lastCommitted == prev {
		cp.nextCP = blockID
	}
	cp.sabft.log.Info("CheckPoint added", commit.ProposedData)
	return nil
}

func (cp *BFTCheckPoint) HasDanglingCheckPoint() bool {
	return cp.NextUncommitted() == nil && len(cp.cache) > 0
}

// (from, to)
// @from is the last committed checkpoint
// @to is any of the dangling uncommitted checkpoints
// Only call it if HasDanglingCheckPoint returns true
func (cp *BFTCheckPoint) MissingRange() (from, to uint64) {
	var v *message.Commit
	for _, v = range cp.cache {
		break
	}
	return cp.lastCommitted.BlockNum(), ExtractBlockID(v).BlockNum()
}

func (cp *BFTCheckPoint) NextUncommitted() *message.Commit {
	if v, ok := cp.cache[cp.lastCommitted]; ok {
		return v
	}
	return nil
}

func (cp *BFTCheckPoint) RemoveNextUncommitted() {
	delete(cp.cache, cp.lastCommitted)
	cp.nextCP = common.EmptyBlockID
}

func (cp *BFTCheckPoint) IsNextCheckPoint(commit *message.Commit) bool {
	id := ExtractBlockID(commit)
	if id == common.EmptyBlockID {
		cp.sabft.log.Fatal("checkpoint on an empty block")
		return false
	}
	_, ok := cp.cache[cp.lastCommitted]
	if !ok {
		return false
	}
	return cp.nextCP == id // && ConvertToBlockID(v.Prev) == cp.lastCommitted
}

func (cp *BFTCheckPoint) Validate(commit *message.Commit) bool {
	// TODO: base validators on last committed block
	//if !cp.sabft.VerifyCommitSig(commit) {
	//	return false
	//}
	return true
}

func (cp *BFTCheckPoint) GetNext(blockNum uint64) (*message.Commit, error) {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, blockNum+1)
	var val []byte
	cp.db.Iterate(key, nil, false, func(key, value []byte) bool {
		val = common.CopyBytes(value)
		return false
	})
	if len(val) == 0 {
		return nil, fmt.Errorf("BFTCheckPoint.GetNext(%d) not found", blockNum)
	}
	commit, err := message.DecodeConsensusMsg(val)
	if err != nil {
		return nil, err
	}
	if err = commit.(*message.Commit).ValidateBasic(); err != nil {
		cp.sabft.log.Error(err)
		return nil, err
	}
	return commit.(*message.Commit), nil
}