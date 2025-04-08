package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AnimateEntity is sent by the server to animate an entity client-side. It may be used to play a single
// animation, or to activate a controller which can start a sequence of animations based on different
// conditions specified in an animation controller.
// Much of the documentation of this packet can be found at
// https://minecraft.gamepedia.com/Bedrock_Edition_beta_animation_documentation.
type AnimateEntity struct {
	Animation            string
	NextState            string
	StopCondition        string
	StopConditionVersion int32
	Controller           string
	BlendOutTime         float32
	EntityRuntimeIDs     []uint64
}

// ID ...
func (*AnimateEntity) ID() uint32 {
	return IDAnimateEntity
}

// Marshal ...
func (pk *AnimateEntity) Marshal(io protocol.IO) {
	io.String(&pk.Animation)
	io.String(&pk.NextState)
	io.String(&pk.StopCondition)
	io.Int32(&pk.StopConditionVersion)
	io.String(&pk.Controller)
	io.Float32(&pk.BlendOutTime)
	protocol.FuncSlice(io, &pk.EntityRuntimeIDs, io.Varuint64)
}
