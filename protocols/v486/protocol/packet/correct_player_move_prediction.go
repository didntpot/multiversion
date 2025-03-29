package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// CorrectPlayerMovePrediction is sent by the server if and only if StartGame.ServerAuthoritativeMovementMode
// is set to AuthoritativeMovementModeServerWithRewind. The packet is used to correct movement at a specific
// point in time.
type CorrectPlayerMovePrediction struct {
	Position mgl32.Vec3
	Delta    mgl32.Vec3
	OnGround bool
	Tick     uint64
}

// ID ...
func (*CorrectPlayerMovePrediction) ID() uint32 {
	return IDCorrectPlayerMovePrediction
}

// Marshal ...
func (pk *CorrectPlayerMovePrediction) Marshal(io protocol.IO) {
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Delta)
	io.Bool(&pk.OnGround)
	io.Varuint64(&pk.Tick)
}
