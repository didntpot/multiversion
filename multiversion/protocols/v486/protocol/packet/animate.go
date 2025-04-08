package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	AnimateActionSwingArm = iota + 1
	_
	AnimateActionStopSleep
	AnimateActionCriticalHit
	AnimateActionMagicCriticalHit
)

const (
	AnimateActionRowRight = iota + 128
	AnimateActionRowLeft
)

// Animate is sent by the server to send a player animation from one player to all viewers of that player. It
// is used for a couple of actions, such as arm swimming and critical hits.
type Animate struct {
	ActionType      int32
	EntityRuntimeID uint64
	BoatRowingTime  float32
}

// ID ...
func (*Animate) ID() uint32 {
	return IDAnimate
}

// Marshal ...
func (pk *Animate) Marshal(io protocol.IO) {
	io.Varint32(&pk.ActionType)
	io.Varuint64(&pk.EntityRuntimeID)
	if pk.ActionType&0x80 != 0 {
		io.Float32(&pk.BoatRowingTime)
	}
}
