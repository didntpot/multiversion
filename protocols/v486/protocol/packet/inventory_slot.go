package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// InventorySlot is sent by the server to update a single slot in one of the inventory windows that the client
// currently has opened. Usually this is the main inventory, but it may also be the off hand or, for example,
// a chest inventory.
type InventorySlot struct {
	WindowID uint32
	Slot     uint32
	NewItem  protocol.ItemInstance
}

// ID ...
func (*InventorySlot) ID() uint32 {
	return IDInventorySlot
}

// Marshal ...
func (pk *InventorySlot) Marshal(io protocol.IO) {
	io.Varuint32(&pk.WindowID)
	io.Varuint32(&pk.Slot)
	io.ItemInstance(&pk.NewItem)
}
