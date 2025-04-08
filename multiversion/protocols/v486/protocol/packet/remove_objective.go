package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// RemoveObjective is sent by the server to remove a scoreboard objective. It is used to stop showing a
// scoreboard to a player.
type RemoveObjective struct {
	ObjectiveName string
}

// ID ...
func (*RemoveObjective) ID() uint32 {
	return IDRemoveObjective
}

// Marshal ...
func (pk *RemoveObjective) Marshal(io protocol.IO) {
	io.String(&pk.ObjectiveName)
}
