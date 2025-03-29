package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	PlayerArmourDamageFlagHelmet = 1 << (iota + 1)
	PlayerArmourDamageFlagChestplate
	PlayerArmourDamageFlagLeggings
	PlayerArmourDamageFlagBoots
)

// PlayerArmourDamage is sent by the server to damage the armour of a player. It is a very efficient packet,
// but generally it's much easier to just send a slot update for the damaged armour.
type PlayerArmourDamage struct {
	Bitset           uint8
	HelmetDamage     int32
	ChestplateDamage int32
	LeggingsDamage   int32
	BootsDamage      int32
}

// ID ...
func (pk *PlayerArmourDamage) ID() uint32 {
	return IDPlayerArmourDamage
}

// Marshal ...
func (pk *PlayerArmourDamage) Marshal(io protocol.IO) {
	io.Uint8(&pk.Bitset)
	if pk.Bitset&PlayerArmourDamageFlagHelmet != 0 {
		io.Varint32(&pk.HelmetDamage)
	} else {
		pk.HelmetDamage = 0
	}
	if pk.Bitset&PlayerArmourDamageFlagChestplate != 0 {
		io.Varint32(&pk.ChestplateDamage)
	} else {
		pk.ChestplateDamage = 0
	}
	if pk.Bitset&PlayerArmourDamageFlagLeggings != 0 {
		io.Varint32(&pk.LeggingsDamage)
	} else {
		pk.LeggingsDamage = 0
	}
	if pk.Bitset&PlayerArmourDamageFlagBoots != 0 {
		io.Varint32(&pk.BootsDamage)
	} else {
		pk.BootsDamage = 0
	}
}
