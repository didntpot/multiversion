package protocol

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ItemComponentEntry is sent in the ItemComponent item table. It holds a name and all the components and
// properties associated to the item.
type ItemComponentEntry struct {
	Name string
	Data map[string]any
}

// Marshal ...
func (x *ItemComponentEntry) Marshal(io protocol.IO) {
	io.String(&x.Name)
	io.NBT(&x.Data, nbt.NetworkLittleEndian)
}
