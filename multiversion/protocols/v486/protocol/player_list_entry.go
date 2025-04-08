package protocol

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PlayerListEntry is an entry found in the PlayerList packet. It represents a single player using the UUID
// found in the entry, and contains several properties such as the skin.
type PlayerListEntry struct {
	protocol.PlayerListEntry
}

// Marshal ...
func (x *PlayerListEntry) Marshal(io protocol.IO) {
	io.UUID(&x.UUID)
	io.Varint64(&x.EntityUniqueID)
	io.String(&x.Username)
	io.String(&x.XUID)
	io.String(&x.PlatformChatID)
	io.Int32(&x.BuildPlatform)
	protocol.Single(io, &Skin{x.Skin})
	io.Bool(&x.Teacher)
	io.Bool(&x.Host)
}
