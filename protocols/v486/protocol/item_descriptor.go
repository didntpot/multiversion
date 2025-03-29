package protocol

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// DefaultItemDescriptor represents an item descriptor for regular items. This is used for the significant majority of
// items.
type DefaultItemDescriptor struct {
	NetworkID     int32
	MetadataValue int32
}

// Marshal ...
func (x *DefaultItemDescriptor) Marshal(io protocol.IO) {
	io.Varint32(&x.NetworkID)
	if x.NetworkID != 0 {
		io.Varint32(&x.MetadataValue)
	}
}
