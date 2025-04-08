package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// MapInfoRequest is sent by the client to request the server to deliver information of a certain map in the
// inventory of the player. The server should respond with a ClientBoundMapItemData packet.
type MapInfoRequest struct {
	MapID int64
}

// ID ...
func (*MapInfoRequest) ID() uint32 {
	return IDMapInfoRequest
}

// Marshal ...
func (pk *MapInfoRequest) Marshal(io protocol.IO) {
	io.Varint64(&pk.MapID)
}
