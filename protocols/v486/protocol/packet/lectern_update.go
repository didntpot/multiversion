package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// LecternUpdate is sent by the client to update the server on which page was opened in a book on a lectern,
// or if the book should be removed from it.
type LecternUpdate struct {
	Page      byte
	PageCount byte
	Position  protocol.BlockPos
	DropBook  bool
}

// ID ...
func (*LecternUpdate) ID() uint32 {
	return IDLecternUpdate
}

// Marshal ...
func (pk *LecternUpdate) Marshal(io protocol.IO) {
	io.Uint8(&pk.Page)
	io.Uint8(&pk.PageCount)
	io.BlockPos(&pk.Position)
	io.Bool(&pk.DropBook)
}
