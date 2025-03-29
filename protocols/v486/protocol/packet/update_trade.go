package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// UpdateTrade is sent by the server to update the trades offered by a villager to a player. It is sent at the
// moment that a player interacts with a villager.
type UpdateTrade struct {
	WindowID          byte
	WindowType        byte
	Size              int32
	TradeTier         int32
	VillagerUniqueID  int64
	EntityUniqueID    int64
	DisplayName       string
	NewTradeUI        bool
	DemandBasedPrices bool
	SerialisedOffers  []byte
}

// ID ...
func (*UpdateTrade) ID() uint32 {
	return IDUpdateTrade
}

// Marshal ...
func (pk *UpdateTrade) Marshal(io protocol.IO) {
	io.Uint8(&pk.WindowID)
	io.Uint8(&pk.WindowType)
	io.Varint32(&pk.Size)
	io.Varint32(&pk.TradeTier)
	io.Varint64(&pk.VillagerUniqueID)
	io.Varint64(&pk.EntityUniqueID)
	io.String(&pk.DisplayName)
	io.Bool(&pk.NewTradeUI)
	io.Bool(&pk.DemandBasedPrices)
	io.Bytes(&pk.SerialisedOffers)
}
