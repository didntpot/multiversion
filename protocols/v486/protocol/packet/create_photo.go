package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// CreatePhoto is a packet that allows players to export photos from their portfolios into items in their inventory.
// This packet only works on the Education Edition version of Minecraft.
type CreatePhoto struct {
	EntityUniqueID int64
	PhotoName      string
	ItemName       string
}

// ID ...
func (*CreatePhoto) ID() uint32 {
	return IDCreatePhoto
}

// Marshal ...
func (pk *CreatePhoto) Marshal(io protocol.IO) {
	io.Int64(&pk.EntityUniqueID)
	io.String(&pk.PhotoName)
	io.String(&pk.ItemName)
}
