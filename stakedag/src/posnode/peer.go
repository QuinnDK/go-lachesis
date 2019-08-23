package posnode

import (
	"time"

	"github.com/Fantom-foundation/go-lachesis/lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/stakedag/src/posnode/api"
)

type (
	// Peer is a representation of other node.
	Peer struct {
		ID     hash.Peer
		Host   string
	}

	// hostAttr contains temporary attributes of host.
	hostAttr struct {
		Name        string
		LastSuccess time.Time
		LastFail    time.Time
	}

	// peerAttr contains temporary attributes of peer.
	peerAttr struct {
		ID   hash.Peer
		Busy bool
		Host *hostAttr
	}
)

// ToWire converts to protobuf message.
func (p *Peer) ToWire() *api.PeerInfo {
	return &api.PeerInfo{
		ID:     p.ID.Hex(),
		Host:   p.Host,
	}
}

// WireToPeer converts from protobuf message.
func WireToPeer(w *api.PeerInfo) *Peer {
	if w == nil {
		return nil
	}
	return &Peer{
		ID:     hash.HexToPeer(w.ID),
		Host:   w.Host,
	}
}

// IDsToWire converts to protobuf message.
func IDsToWire(ids []hash.Peer) *api.PeerIDs {
	w := &api.PeerIDs{
		IDs: make([]string, len(ids)),
	}

	for i, id := range ids {
		w.IDs[i] = id.Hex()
	}

	return w
}

// WireToIDs converts from protobuf message.
func WireToIDs(w *api.PeerIDs) []hash.Peer {
	if w == nil {
		return nil
	}

	res := make([]hash.Peer, len(w.IDs))
	for i, str := range w.IDs {
		res[i] = hash.HexToPeer(str)
	}

	return res
}
