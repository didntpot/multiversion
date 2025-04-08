package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Emote is sent by both the server and the client. When the client sends an emote, it sends this packet to
// the server, after which the server will broadcast the packet to other players online.
type Emote struct {
	EntityRuntimeID uint64
	EmoteID         string
	Flags           byte
}

// ID ...
func (*Emote) ID() uint32 {
	return IDEmote
}

// Marshal ...
func (pk *Emote) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.String(&pk.EmoteID)
	io.Uint8(&pk.Flags)
}
