package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// MobEquipment is sent by the client to the server and the server to the client to make the other side
// aware of the new item that an entity is holding. It is used to show the item in the hand of entities such
// as zombies too.
type MobEquipment struct {
	EntityRuntimeID uint64
	NewItem         protocol.ItemInstance
	InventorySlot   byte
	HotBarSlot      byte
	WindowID        byte
}

// ID ...
func (*MobEquipment) ID() uint32 {
	return IDMobEquipment
}

// Marshal ...
func (pk *MobEquipment) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.ItemInstance(&pk.NewItem)
	io.Uint8(&pk.InventorySlot)
	io.Uint8(&pk.HotBarSlot)
	io.Uint8(&pk.WindowID)
}
