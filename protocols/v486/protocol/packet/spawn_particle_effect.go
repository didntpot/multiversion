package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SpawnParticleEffect is sent by the server to spawn a particle effect client-side. Unlike other packets that
// result in the appearing of particles, this packet can show particles that are not hardcoded in the client.
// They can be added and changed through behaviour packs to implement custom particles.
type SpawnParticleEffect struct {
	Dimension      byte
	EntityUniqueID int64
	Position       mgl32.Vec3
	ParticleName   string
}

// ID ...
func (*SpawnParticleEffect) ID() uint32 {
	return IDSpawnParticleEffect
}

// Marshal ...
func (pk *SpawnParticleEffect) Marshal(io protocol.IO) {
	io.Uint8(&pk.Dimension)
	io.Varint64(&pk.EntityUniqueID)
	io.Vec3(&pk.Position)
	io.String(&pk.ParticleName)
}
