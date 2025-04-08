package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SubChunkRequest requests specific sub-chunks from the server using a center point.
type SubChunkRequest struct {
	Dimension int32
	Position  protocol.SubChunkPos
	Offsets   []protocol.SubChunkOffset
}

// ID ...
func (*SubChunkRequest) ID() uint32 {
	return IDSubChunkRequest
}

// Marshal ...
func (pk *SubChunkRequest) Marshal(io protocol.IO) {
	io.Varint32(&pk.Dimension)
	io.SubChunkPos(&pk.Position)
	protocol.SliceUint32Length(io, &pk.Offsets)
}
