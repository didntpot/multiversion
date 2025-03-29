package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	GameTypeSurvival = iota
	GameTypeCreative
	GameTypeAdventure
	GameTypeSurvivalSpectator
	GameTypeCreativeSpectator
)

// SetPlayerGameType is sent by the server to update the game type, which is otherwise known as the game mode,
// of a player.
type SetPlayerGameType struct {
	GameType int32
}

// ID ...
func (*SetPlayerGameType) ID() uint32 {
	return IDSetPlayerGameType
}

// Marshal ...
func (pk *SetPlayerGameType) Marshal(io protocol.IO) {
	io.Varint32(&pk.GameType)
}
