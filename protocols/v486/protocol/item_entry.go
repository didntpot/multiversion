package protocol

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ItemEntry is an item sent in the StartGame item table. It holds a name and a legacy ID, which is used to
// point back to that name.
type ItemEntry struct {
	Name           string
	RuntimeID      int16
	ComponentBased bool
}

// Marshal ...
func (x *ItemEntry) Marshal(io protocol.IO) {
	io.String(&x.Name)
	io.Int16(&x.RuntimeID)
	io.Bool(&x.ComponentBased)
}
