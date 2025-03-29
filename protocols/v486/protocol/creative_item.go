package protocol

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// CreativeItem represents a creative item present in the creative inventory.
type CreativeItem struct {
	CreativeItemNetworkID uint32
	Item                  protocol.ItemStack
}

// Marshal ...
func (x *CreativeItem) Marshal(io protocol.IO) {
	io.Varuint32(&x.CreativeItemNetworkID)
	io.Item(&x.Item)
}
