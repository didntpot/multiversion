package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	ShowCreditsStatusStart = iota
	ShowCreditsStatusEnd
)

// ShowCredits is sent by the server to show the Minecraft credits screen to the client. It is typically sent
// when the player beats the ender dragon and leaves the End.
type ShowCredits struct {
	PlayerRuntimeID uint64
	StatusType      int32
}

// ID ...
func (*ShowCredits) ID() uint32 {
	return IDShowCredits
}

// Marshal ...
func (pk *ShowCredits) Marshal(io protocol.IO) {
	io.Varuint64(&pk.PlayerRuntimeID)
	io.Varint32(&pk.StatusType)
}
