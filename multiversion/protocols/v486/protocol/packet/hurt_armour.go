package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// HurtArmour is sent by the server to damage the player's armour after being hit. The packet should never be
// used by servers as it hands the responsibility over to the player completely, while the server can easily
// reliably update the armour damage of players itself.
type HurtArmour struct {
	Cause       int32
	Damage      int32
	ArmourSlots int64
}

// ID ...
func (*HurtArmour) ID() uint32 {
	return IDHurtArmour
}

// Marshal ...
func (pk *HurtArmour) Marshal(io protocol.IO) {
	io.Varint32(&pk.Cause)
	io.Varint32(&pk.Damage)
	io.Varint64(&pk.ArmourSlots)
}
