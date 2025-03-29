package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// StopSound is sent by the server to stop a sound playing to the player, such as a playing music disk track
// or other long-lasting sounds.
type StopSound struct {
	SoundName string
	StopAll   bool
}

// ID ...
func (*StopSound) ID() uint32 {
	return IDStopSound
}

// Marshal ...
func (pk *StopSound) Marshal(io protocol.IO) {
	io.String(&pk.SoundName)
	io.Bool(&pk.StopAll)
}
