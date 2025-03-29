package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/protocols/v486/protocol"
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PlayerSkin is sent by the client to the server when it updates its own skin using the in-game skin picker.
// It is relayed by the server, or sent if the server changes the skin of a player on its own accord. Note
// that the packet can only be sent for players that are in the player list at the time of sending.
type PlayerSkin struct {
	UUID        uuid.UUID
	Skin        legacyprotocol.Skin
	NewSkinName string
	OldSkinName string
}

// ID ...
func (*PlayerSkin) ID() uint32 {
	return IDPlayerSkin
}

// Marshal ...
func (pk *PlayerSkin) Marshal(io protocol.IO) {
	io.UUID(&pk.UUID)
	protocol.Single(io, &pk.Skin)
	io.String(&pk.NewSkinName)
	io.String(&pk.OldSkinName)
	io.Bool(&pk.Skin.Trusted)
}
