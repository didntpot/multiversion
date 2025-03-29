package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// UpdateEquip is sent by the server to the client upon opening a horse inventory. It is used to set the
// content of the inventory and specify additional properties, such as the items that are allowed to be put
// in slots of the inventory.
type UpdateEquip struct {
	WindowID                byte
	WindowType              byte
	Size                    int32
	EntityUniqueID          int64
	SerialisedInventoryData []byte
}

// ID ...
func (*UpdateEquip) ID() uint32 {
	return IDUpdateEquip
}

// Marshal ...
func (pk *UpdateEquip) Marshal(io protocol.IO) {
	io.Uint8(&pk.WindowID)
	io.Uint8(&pk.WindowType)
	io.Varint32(&pk.Size)
	io.Varint64(&pk.EntityUniqueID)
	io.Bytes(&pk.SerialisedInventoryData)
}
