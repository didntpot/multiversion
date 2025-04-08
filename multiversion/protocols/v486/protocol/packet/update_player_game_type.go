package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// UpdatePlayerGameType is sent by the server to change the game mode of a player. It is functionally
// identical to the SetPlayerGameType packet.
type UpdatePlayerGameType struct {
	GameType       int32
	PlayerUniqueID int64
}

// ID ...
func (*UpdatePlayerGameType) ID() uint32 {
	return IDUpdatePlayerGameType
}

// Marshal ...
func (pk *UpdatePlayerGameType) Marshal(io protocol.IO) {
	io.Varint32(&pk.GameType)
	io.Varint64(&pk.PlayerUniqueID)
}
