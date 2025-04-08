package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PlaySound is sent by the server to play a sound to the client. Some of the sounds may only be started using
// this packet and must be stopped using the StopSound packet.
type PlaySound struct {
	SoundName string
	Position  mgl32.Vec3
	Volume    float32
	Pitch     float32
}

// ID ...
func (*PlaySound) ID() uint32 {
	return IDPlaySound
}

// Marshal ...
func (pk *PlaySound) Marshal(io protocol.IO) {
	b := protocol.BlockPos{int32(pk.Position[0] * 8), int32(pk.Position[1] * 8), int32(pk.Position[2] * 8)}

	io.String(&pk.SoundName)
	io.BlockPos(&b)
	io.Float32(&pk.Volume)
	io.Float32(&pk.Pitch)
}
