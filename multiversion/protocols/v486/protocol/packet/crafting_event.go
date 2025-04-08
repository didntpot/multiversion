package packet

import (
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// CraftingEvent is sent by the client when it crafts a particular item. Note that this packet may be fully
// ignored, as the InventoryTransaction packet provides all the information required.
type CraftingEvent struct {
	WindowID     byte
	CraftingType int32
	RecipeUUID   uuid.UUID
	Input        []protocol.ItemInstance
	Output       []protocol.ItemInstance
}

// ID ...
func (*CraftingEvent) ID() uint32 {
	return IDCraftingEvent
}

// Marshal ...
func (pk *CraftingEvent) Marshal(io protocol.IO) {
	inputLen, outputLen := uint32(len(pk.Input)), uint32(len(pk.Output))
	io.Uint8(&pk.WindowID)
	io.Varint32(&pk.CraftingType)
	io.UUID(&pk.RecipeUUID)
	io.Varuint32(&inputLen)
	for _, input := range pk.Input {
		io.ItemInstance(&input)
	}
	io.Varuint32(&outputLen)
	for _, output := range pk.Output {
		io.ItemInstance(&output)
	}
}
