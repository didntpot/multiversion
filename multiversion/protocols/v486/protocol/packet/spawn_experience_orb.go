package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SpawnExperienceOrb is sent by the server to spawn an experience orb entity client-side. Much like the
// AddPainting packet, it is one of the few packets that spawn an entity without using the AddActor packet.
type SpawnExperienceOrb struct {
	Position         mgl32.Vec3
	ExperienceAmount int32
}

// ID ...
func (*SpawnExperienceOrb) ID() uint32 {
	return IDSpawnExperienceOrb
}

// Marshal ...
func (pk *SpawnExperienceOrb) Marshal(io protocol.IO) {
	io.Vec3(&pk.Position)
	io.Varint32(&pk.ExperienceAmount)
}
