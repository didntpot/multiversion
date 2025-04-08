package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// MobEffect is sent by the server to apply an effect to the player, for example an effect like poison. It may
// also be used to modify existing effects, or removing them completely.
type MobEffect struct {
	EntityRuntimeID uint64
	Operation       byte
	EffectType      int32
	Amplifier       int32
	Particles       bool
	Duration        int32
}

// ID ...
func (*MobEffect) ID() uint32 {
	return IDMobEffect
}

// Marshal ...
func (pk *MobEffect) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Uint8(&pk.Operation)
	io.Varint32(&pk.EffectType)
	io.Varint32(&pk.Amplifier)
	io.Bool(&pk.Particles)
	io.Varint32(&pk.Duration)
}
