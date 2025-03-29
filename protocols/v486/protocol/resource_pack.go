package protocol

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// BehaviourPackInfo represents a behaviour pack's info sent over network. It holds information about the
// behaviour pack such as its name, description and version.
type BehaviourPackInfo struct {
	UUID            string
	Version         string
	Size            uint64
	ContentKey      string
	SubPackName     string
	ContentIdentity string
	HasScripts      bool
}

// Marshal ...
func (x *BehaviourPackInfo) Marshal(io protocol.IO) {
	io.String(&x.UUID)
	io.String(&x.Version)
	io.Uint64(&x.Size)
	io.String(&x.ContentKey)
	io.String(&x.SubPackName)
	io.String(&x.ContentIdentity)
	io.Bool(&x.HasScripts)
}

// TexturePackInfo represents a texture pack's info sent over network. It holds information about the
// texture pack such as its name, description and version.
type TexturePackInfo struct {
	UUID            string
	Version         string
	Size            uint64
	ContentKey      string
	SubPackName     string
	ContentIdentity string
	HasScripts      bool
	RTXEnabled      bool
}

// Marshal ...
func (x *TexturePackInfo) Marshal(io protocol.IO) {
	io.String(&x.UUID)
	io.String(&x.Version)
	io.Uint64(&x.Size)
	io.String(&x.ContentKey)
	io.String(&x.SubPackName)
	io.String(&x.ContentIdentity)
	io.Bool(&x.HasScripts)
	io.Bool(&x.RTXEnabled)
}

// StackResourcePack represents a resource pack sent on the stack of the client. When sent, the client will
// apply them in the order of the stack sent.
type StackResourcePack struct {
	UUID        string
	Version     string
	SubPackName string
}

// Marshal ...
func (x *StackResourcePack) Marshal(io protocol.IO) {
	io.String(&x.UUID)
	io.String(&x.Version)
	io.String(&x.SubPackName)
}

// PackURL represents a resource pack that is being served from a HTTP server rather than being sent over
// the Minecraft protocol.
type PackURL struct {
	UUIDVersion string
	URL         string
}

// Marshal ...
func (x *PackURL) Marshal(io protocol.IO) {
	io.String(&x.UUIDVersion)
	io.String(&x.URL)
}
