package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// MobArmourEquipment is sent by the server to the client to update the armour an entity is wearing. It is
// sent for both players and other entities, such as zombies.
type MobArmourEquipment struct {
	EntityRuntimeID uint64
	Helmet          protocol.ItemInstance
	Chestplate      protocol.ItemInstance
	Leggings        protocol.ItemInstance
	Boots           protocol.ItemInstance
}

// ID ...
func (*MobArmourEquipment) ID() uint32 {
	return IDMobArmourEquipment
}

// Marshal ...
func (pk *MobArmourEquipment) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.ItemInstance(&pk.Helmet)
	io.ItemInstance(&pk.Chestplate)
	io.ItemInstance(&pk.Leggings)
	io.ItemInstance(&pk.Boots)
}
