package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

const (
	CameraShakeTypePositional uint8 = iota
	CameraShakeTypeRotational
)

const (
	CameraShakeActionAdd = iota
	CameraShakeActionStop
)

// CameraShake is sent by the server to make the camera shake client-side. This feature was added for map-
// making partners.
type CameraShake struct {
	Intensity float32
	Duration  float32
	Type      uint8
	Action    uint8
}

// ID ...
func (*CameraShake) ID() uint32 {
	return IDCameraShake
}

// Marshal ...
func (pk *CameraShake) Marshal(io protocol.IO) {
	io.Float32(&pk.Intensity)
	io.Float32(&pk.Duration)
	io.Uint8(&pk.Type)
	io.Uint8(&pk.Action)
}
