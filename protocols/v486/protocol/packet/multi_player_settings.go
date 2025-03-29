package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	EnableMultiPlayer = iota
	DisableMultiPlayer
	RefreshJoinCode
)

// MultiPlayerSettings is sent by the client to update multi-player related settings server-side and sent back
// to online players by the server.
// The MultiPlayerSettings packet is a Minecraft: Education Edition packet. It has no functionality for the
// base game.
type MultiPlayerSettings struct {
	ActionType int32
}

// ID ...
func (*MultiPlayerSettings) ID() uint32 {
	return IDMultiPlayerSettings
}

// Marshal ...
func (pk *MultiPlayerSettings) Marshal(io protocol.IO) {
	io.Varint32(&pk.ActionType)
}
