// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"errors"
	"time"

	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/peering/domain"
	"github.com/iotaledger/wasp/packages/peering/group"
	"golang.org/x/xerrors"
)

//
// PeeringNetwork represents a global view of the mocked network.
//
type PeeringNetwork struct {
	nodes     []*peeringNode
	providers []*peeringNetworkProvider
	bufSize   int
	behavior  PeeringNetBehavior
	log       *logger.Logger
}

// NewPeeringNetworkForLocs creates a test network with new keys, etc.
func NewPeeringNetworkForLocs(peerNetIDs []string, bufSize int, log *logger.Logger) *PeeringNetwork {
	identities := make([]*ed25519.KeyPair, len(peerNetIDs))
	for i := range identities {
		nodeIdentity := ed25519.GenerateKeyPair()
		identities[i] = &nodeIdentity
	}
	behavior := NewPeeringNetReliable(log)
	return NewPeeringNetwork(peerNetIDs, identities, bufSize, behavior, log)
}

// NewPeeringNetwork creates new test network, it can then be used to create network nodes.
func NewPeeringNetwork(
	netIDs []string,
	nodeIdentities []*ed25519.KeyPair,
	bufSize int,
	behavior PeeringNetBehavior,
	log *logger.Logger,
) *PeeringNetwork {
	nodes := make([]*peeringNode, len(netIDs))
	providers := make([]*peeringNetworkProvider, len(netIDs))
	network := PeeringNetwork{
		nodes:     nodes,
		providers: providers,
		bufSize:   bufSize,
		behavior:  behavior,
		log:       log,
	}
	for i := range nodes {
		nodes[i] = newPeeringNode(netIDs[i], nodeIdentities[i], &network)
	}
	for i := range nodes {
		providers[i] = newPeeringNetworkProvider(nodes[i], &network)
	}
	return &network
}

// NetworkProviders returns network providers for each of the nodes in the network.
func (p *PeeringNetwork) NetworkProviders() []peering.NetworkProvider {
	cp := make([]peering.NetworkProvider, len(p.providers))
	for i := range p.providers {
		cp[i] = p.providers[i]
	}
	return cp
}

func (p *PeeringNetwork) nodeByNetID(nodeNetID string) *peeringNode {
	for i := range p.nodes {
		if p.nodes[i].netID == nodeNetID {
			return p.nodes[i]
		}
	}
	return nil
}

// Close implements the io.Closer interface.
func (p *PeeringNetwork) Close() error {
	for _, n := range p.nodes {
		if err := n.Close(); err != nil {
			return xerrors.Errorf("failed to close test peering node: %w", err)
		}
	}
	p.behavior.Close()
	return nil
}

//
// peeringNode stands for a mock of a node in a fake network.
// It does NOT implement the peering.PeerSender, because the source
// node should be known for the sender.
//
type peeringNode struct {
	netID    string
	identity *ed25519.KeyPair
	sendCh   chan *peeringMsg
	recvCh   chan *peeringMsg
	recvCbs  []*peeringCb
	network  *PeeringNetwork
	log      *logger.Logger
}

type peeringMsg struct {
	from *peeringNode
	msg  peering.PeerMessage
}

type peeringCb struct {
	callback  func(recv *peering.RecvEvent) // Receive callback.
	destNP    *peeringNetworkProvider       // Destination node.
	peeringID *peering.PeeringID            // Only listen for specific chain msgs.
}

func newPeeringNode(netID string, identity *ed25519.KeyPair, network *PeeringNetwork) *peeringNode {
	sendCh := make(chan *peeringMsg, network.bufSize)
	recvCh := make(chan *peeringMsg, network.bufSize)
	recvCbs := make([]*peeringCb, 0)
	n := peeringNode{
		netID:    netID,
		identity: identity,
		sendCh:   sendCh,
		recvCh:   recvCh,
		recvCbs:  recvCbs,
		network:  network,
		log:      network.log.With("loc", netID),
	}
	network.behavior.AddLink(sendCh, recvCh, netID)
	go n.recvLoop()
	return &n
}

func (n *peeringNode) recvLoop() {
	for pm := range n.recvCh {
		msgPeeringID := pm.msg.PeeringID.String()
		for _, cb := range n.recvCbs {
			if cb.peeringID == nil || cb.peeringID.String() == msgPeeringID {
				cb.callback(&peering.RecvEvent{
					From: cb.destNP.senderByNetID(pm.from.netID),
					Msg:  &pm.msg,
				})
			}
		}
	}
}

func (n *peeringNode) sendMsg(from *peeringNode, msg *peering.PeerMessage) {
	n.sendCh <- &peeringMsg{
		from: from,
		msg:  *msg,
	}
}

func (n *peeringNode) Close() error {
	close(n.recvCh)
	return nil
}

//
// peeringNetworkProvider to be used in tests as a mock for the peering network.
//
type peeringNetworkProvider struct {
	self    *peeringNode
	network *PeeringNetwork
	senders []*peeringSender // Senders for all the nodes.
}

// NewpeeringNetworkProvider initializes new network provider (a local view).
func newPeeringNetworkProvider(self *peeringNode, network *PeeringNetwork) *peeringNetworkProvider {
	senders := make([]*peeringSender, len(network.nodes))
	netProvider := peeringNetworkProvider{
		self:    self,
		network: network,
		senders: senders,
	}
	for i := range network.nodes {
		senders[i] = newPeeringSender(network.nodes[i], &netProvider)
	}
	return &netProvider
}

// Run implements peering.NetworkProvider.
func (p *peeringNetworkProvider) Run(stopCh <-chan struct{}) {
	<-stopCh
}

// Self implements peering.NetworkProvider.
func (p *peeringNetworkProvider) Self() peering.PeerSender {
	return newPeeringSender(p.self, p)
}

// Group implements peering.NetworkProvider.
func (p *peeringNetworkProvider) PeerGroup(peerAddrs []string) (peering.GroupProvider, error) {
	peers := make([]peering.PeerSender, len(peerAddrs))
	for i := range peerAddrs {
		n := p.network.nodeByNetID(peerAddrs[i])
		if n == nil {
			return nil, errors.New("unknown_node_location")
		}
		peers[i] = p.senders[i]
	}
	return group.NewPeeringGroupProvider(p, peers, p.network.log)
}

// Domain creates peering.PeerDomainProvider.
func (p *peeringNetworkProvider) PeerDomain(peerNetIDs []string) (peering.PeerDomainProvider, error) {
	return domain.NewPeerDomainByNetIDs(p, peerNetIDs, p.network.log)
}

// Attach implements peering.NetworkProvider.
func (p *peeringNetworkProvider) Attach(
	peeringID *peering.PeeringID,
	callback func(recv *peering.RecvEvent),
) interface{} {
	p.self.recvCbs = append(p.self.recvCbs, &peeringCb{
		callback:  callback,
		destNP:    p,
		peeringID: peeringID,
	})
	return nil // We don't care on the attachIDs for now.
}

// Detach implements peering.NetworkProvider.
func (p *peeringNetworkProvider) Detach(attachID interface{}) {
	// Detach is not important in tests.
}

// PeerByNetID implements peering.NetworkProvider.
func (p *peeringNetworkProvider) PeerByNetID(peerNetID string) (peering.PeerSender, error) {
	if s := p.senderByNetID(peerNetID); s != nil {
		return s, nil
	}
	return nil, errors.New("peer not found by NetID")
}

// PeerByNetID implements peering.NetworkProvider.
func (p *peeringNetworkProvider) PeerByPubKey(peerPub *ed25519.PublicKey) (peering.PeerSender, error) {
	for i := range p.senders {
		if p.senders[i].node.identity.PublicKey == *peerPub {
			return p.senders[i], nil
		}
	}
	return nil, errors.New("peer not found by pubKey")
}

// PeerStatus implements peering.NetworkProvider.
func (p *peeringNetworkProvider) PeerStatus() []peering.PeerStatusProvider {
	peerStatus := make([]peering.PeerStatusProvider, len(p.senders))
	for i := range peerStatus {
		peerStatus[i] = p.senders[i]
	}
	return peerStatus
}

func (p *peeringNetworkProvider) senderByNetID(peerNetID string) *peeringSender {
	for i := range p.senders {
		if p.senders[i].node.netID == peerNetID {
			return p.senders[i]
		}
	}
	return nil
}

//
// peeringSender represents a local view of a remote node
// and implements the peering.PeerSender interface.
//
type peeringSender struct {
	node        *peeringNode
	netProvider *peeringNetworkProvider
}

func newPeeringSender(node *peeringNode, netProvider *peeringNetworkProvider) *peeringSender {
	return &peeringSender{
		node:        node,
		netProvider: netProvider,
	}
}

// NetID implements peering.PeerSender.
func (p *peeringSender) NetID() string {
	return p.node.netID
}

// PubKey implements peering.PeerSender.
func (p *peeringSender) PubKey() *ed25519.PublicKey {
	return &p.node.identity.PublicKey
}

// Send implements peering.PeerSender.
func (p *peeringSender) SendMsg(msg *peering.PeerMessage) {
	p.node.sendMsg(p.netProvider.self, msg)
}

// IsAlive implements peering.PeerSender.
func (p *peeringSender) IsAlive() bool {
	return true // Not needed in tests.
}

// Await implements peering.PeerSender.
func (p *peeringSender) Await(timeout time.Duration) error {
	return nil
}

// IsInbound implements peering.PeerStatusProvider.
func (p *peeringSender) IsInbound() bool {
	return true // Not needed in tests.
}

// NumUsers implements peering.PeerStatusProvider.
func (p *peeringSender) NumUsers() int {
	return 0 // Not needed in tests.
}

// Send implements peering.PeerSender.
func (p *peeringSender) Close() {
	// Not needed in tests.
}
