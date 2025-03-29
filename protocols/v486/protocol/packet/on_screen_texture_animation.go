package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// OnScreenTextureAnimation is sent by the server to show a certain animation on the screen of the player.
// The packet is used, as an example, for when a raid is triggered and when a raid is defeated.
type OnScreenTextureAnimation struct {
	AnimationType int32
}

// ID ...
func (*OnScreenTextureAnimation) ID() uint32 {
	return IDOnScreenTextureAnimation
}

// Marshal ...
func (pk *OnScreenTextureAnimation) Marshal(io protocol.IO) {
	io.Int32(&pk.AnimationType)
}
