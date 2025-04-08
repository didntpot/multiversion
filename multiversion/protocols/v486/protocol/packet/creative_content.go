package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// CreativeContent is a packet sent by the server to set the creative inventory's content for a player.
// Introduced in 1.16, this packet replaces the previous method - sending an InventoryContent packet with
// creative inventory window ID.
// From v1.16.100 until 1.21.5x, this packet must be sent during the login sequence. Not sending it will stop the client
// from joining the server.
type CreativeContent struct {
	Items []legacyprotocol.CreativeItem
}

// ID ...
func (*CreativeContent) ID() uint32 {
	return IDCreativeContent
}

// Marshal ...
func (pk *CreativeContent) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Items)
}
