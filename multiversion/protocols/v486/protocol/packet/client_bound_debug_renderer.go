package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	ClientBoundDebugRendererClear int32 = iota + 1
	ClientBoundDebugRendererAddCube
)

// ClientBoundDebugRenderer is sent by the server to spawn an outlined cube on client-side.
type ClientBoundDebugRenderer struct {
	Type     int32
	Text     string
	Position mgl32.Vec3
	Red      float32
	Green    float32
	Blue     float32
	Alpha    float32
	Duration int64
}

// ID ...
func (*ClientBoundDebugRenderer) ID() uint32 {
	return IDClientBoundDebugRenderer
}

// Marshal ...
func (pk *ClientBoundDebugRenderer) Marshal(io protocol.IO) {
	io.Int32(&pk.Type)
	if pk.Type == ClientBoundDebugRendererAddCube {
		io.String(&pk.Text)
		io.Vec3(&pk.Position)
		io.Float32(&pk.Red)
		io.Float32(&pk.Green)
		io.Float32(&pk.Blue)
		io.Float32(&pk.Alpha)
		io.Int64(&pk.Duration)
	}
}
