package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ResourcePackStack is sent by the server to send the order in which resource packs and behaviour packs
// should be applied (and downloaded) by the client.
type ResourcePackStack struct {
	TexturePackRequired          bool
	BehaviourPacks               []protocol.StackResourcePack
	TexturePacks                 []protocol.StackResourcePack
	BaseGameVersion              string
	Experiments                  []protocol.ExperimentData
	ExperimentsPreviouslyToggled bool
}

// ID ...
func (*ResourcePackStack) ID() uint32 {
	return IDResourcePackStack
}

// Marshal ...
func (pk *ResourcePackStack) Marshal(io protocol.IO) {
	io.Bool(&pk.TexturePackRequired)
	protocol.Slice(io, &pk.BehaviourPacks)
	protocol.Slice(io, &pk.TexturePacks)
	io.String(&pk.BaseGameVersion)
	protocol.SliceUint32Length(io, &pk.Experiments)
	io.Bool(&pk.ExperimentsPreviouslyToggled)
}
