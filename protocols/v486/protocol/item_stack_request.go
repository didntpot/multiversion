package protocol

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// ItemStackRequest represents a single request present in an ItemStackRequest packet sent by the client to
// change an item in an inventory.
// Item stack requests are either approved or rejected by the server using the ItemStackResponse packet.
type ItemStackRequest struct {
	protocol.ItemStackRequest
}

// Marshal ...
func (x *ItemStackRequest) Marshal(io protocol.IO) {
	io.Varint32(&x.RequestID)
	protocol.FuncSlice(io, &x.Actions, io.StackRequestAction)
	protocol.FuncSlice(io, &x.FilterStrings, io.String)
}
