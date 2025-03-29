package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// MotionPredictionHints is sent by the server to the client. There is a predictive movement component for
// entities. This packet fills the "history" of that component and entity movement is computed based on the
// points. Vanilla sends this packet instead of the SetActorMotion packet when 'spatial optimisations' are
// enabled.
type MotionPredictionHints struct {
	EntityRuntimeID uint64
	Velocity        mgl32.Vec3
	OnGround        bool
}

// ID ...
func (*MotionPredictionHints) ID() uint32 {
	return IDMotionPredictionHints
}

// Marshal ...
func (pk *MotionPredictionHints) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Vec3(&pk.Velocity)
	io.Bool(&pk.OnGround)
}
