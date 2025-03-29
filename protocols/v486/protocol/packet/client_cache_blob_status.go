package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ClientCacheBlobStatus is part of the blob cache protocol. It is sent by the client to let the server know
// what blobs it needs and which blobs it already has, in an ACK type system.
type ClientCacheBlobStatus struct {
	MissHashes []uint64
	HitHashes  []uint64
}

// ID ...
func (pk *ClientCacheBlobStatus) ID() uint32 {
	return IDClientCacheBlobStatus
}

// Marshal ...
func (pk *ClientCacheBlobStatus) Marshal(io protocol.IO) {
	missLen, hitLen := uint32(len(pk.MissHashes)), uint32(len(pk.HitHashes))
	io.Varuint32(&missLen)
	io.Varuint32(&hitLen)
	protocol.FuncSliceOfLen(io, missLen, &pk.MissHashes, io.Uint64)
	protocol.FuncSliceOfLen(io, hitLen, &pk.HitHashes, io.Uint64)
}
