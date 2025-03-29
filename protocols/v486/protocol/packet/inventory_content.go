package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// InventoryContent is sent by the server to update the full content of a particular inventory. It is usually
// sent for the main inventory of the player, but also works for other inventories that are currently opened
// by the player.
type InventoryContent struct {
	WindowID uint32
	Content  []protocol.ItemInstance
}

// ID ...
func (*InventoryContent) ID() uint32 {
	return IDInventoryContent
}

// Marshal ...
func (pk *InventoryContent) Marshal(io protocol.IO) {
	io.Varuint32(&pk.WindowID)
	protocol.FuncSlice(io, &pk.Content, io.ItemInstance)
}
