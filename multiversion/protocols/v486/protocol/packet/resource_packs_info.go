package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ResourcePacksInfo is sent by the server to inform the client on what resource packs the server has. It
// sends a list of the resource packs it has and basic information on them like the version and description.
type ResourcePacksInfo struct {
	TexturePackRequired bool
	HasScripts          bool
	BehaviourPacks      []legacyprotocol.BehaviourPackInfo
	TexturePacks        []legacyprotocol.TexturePackInfo
	ForcingServerPacks  bool
}

// ID ...
func (*ResourcePacksInfo) ID() uint32 {
	return IDResourcePacksInfo
}

// Marshal ...
func (pk *ResourcePacksInfo) Marshal(io protocol.IO) {
	io.Bool(&pk.TexturePackRequired)
	io.Bool(&pk.HasScripts)
	io.Bool(&pk.ForcingServerPacks)
	protocol.SliceUint16Length(io, &pk.BehaviourPacks)
	protocol.SliceUint16Length(io, &pk.TexturePacks)
}
