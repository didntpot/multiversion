package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	_ = iota + 1
	_
	InteractActionLeaveVehicle
	InteractActionMouseOverEntity
	InteractActionNPCOpen
	InteractActionOpenInventory
)

// Interact is sent by the client when it interacts with another entity in some way. It used to be used for
// normal entity and block interaction, but this is no longer the case now.
type Interact struct {
	ActionType            byte
	TargetEntityRuntimeID uint64
	Position              mgl32.Vec3
}

// ID ...
func (*Interact) ID() uint32 {
	return IDInteract
}

// Marshal ...
func (pk *Interact) Marshal(io protocol.IO) {
	io.Uint8(&pk.ActionType)
	io.Varuint64(&pk.TargetEntityRuntimeID)
	switch pk.ActionType {
	case InteractActionMouseOverEntity, InteractActionLeaveVehicle:
		io.Vec3(&pk.Position)
	}
}
