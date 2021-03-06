package p2p

import (
	"github.com/coschain/contentos-go/node"
	"github.com/coschain/contentos-go/p2p/message/types"
	"github.com/coschain/contentos-go/p2p/peer"
	"github.com/sirupsen/logrus"
)

//P2P represent the net interface of p2p package
type P2P interface {
	Start()
	Halt()
	Connect(addr string, isConsensus bool) error
	GetID() uint64
	GetVersion() uint32
	GetSyncPort() uint32
	GetConsPort() uint32
	GetRelay() bool
	GetHeight() uint64
	GetTime() int64
	GetServices() uint64
	GetNeighbors() []*peer.Peer
	GetNeighborAddrs() []*types.PeerAddr
	GetConnectionCnt() uint32
	GetNp() *peer.NbrPeers
	GetPeer(uint64) *peer.Peer
	SetHeight(uint64)
	IsPeerEstablished(p *peer.Peer) bool
	GetMsgChan(isConsensus bool) chan *types.MsgPayload
	GetPeerFromAddr(addr string) *peer.Peer
	AddOutConnectingList(addr string) (added bool)
	GetOutConnRecordLen() int
	RemoveFromConnectingList(addr string)
	RemoveFromOutConnRecord(addr string)
	RemoveFromInConnRecord(addr string)
	AddPeerSyncAddress(addr string, p *peer.Peer)
	AddPeerConsAddress(addr string, p *peer.Peer)
	GetOutConnectingListLen() (count uint)
	RemovePeerSyncAddress(addr string)
	RemovePeerConsAddress(addr string)
	AddNbrNode(*peer.Peer)
	DelNbrNode(p *peer.Peer) (*peer.Peer, bool)
	NodeEstablished(uint64) bool
	SetOwnAddress(addr string)
	IsAddrFromConnecting(addr string) bool
	Broadcast(msg types.Message, isConsensus bool)
	Send(p *peer.Peer, msg types.Message, isConsensus bool) error

	GetService(string) (interface{}, error)
	GetContex() *node.ServiceContext
	GetLog() *logrus.Logger
}
