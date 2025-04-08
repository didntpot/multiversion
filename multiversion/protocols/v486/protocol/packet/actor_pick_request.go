package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ActorPickRequest is sent by the client when it tries to pick an entity, so that it gets a spawn egg which
// can spawn that entity.
type ActorPickRequest struct {
	EntityUniqueID int64
	HotBarSlot     byte
	WithData       bool
}

// ID ...
func (*ActorPickRequest) ID() uint32 {
	return IDActorPickRequest
}

// Marshal ...
func (pk *ActorPickRequest) Marshal(io protocol.IO) {
	io.Int64(&pk.EntityUniqueID)
	io.Uint8(&pk.HotBarSlot)
	io.Bool(&pk.WithData)
}
