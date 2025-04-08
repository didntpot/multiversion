package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ShowStoreOffer is sent by the server to show a Marketplace store offer to a player. It opens a window
// client-side that displays the item.
// The ShowStoreOffer packet only works on the partnered servers: Servers that are not partnered will not have
// a store buttons show up in the in-game pause menu and will, as a result, not be able to open store offers
// on the client side. Sending the packet does therefore not work when using a proxy that is not connected to
// with the domain of one of the partnered servers.
type ShowStoreOffer struct {
	OfferID string
	ShowAll bool
}

// ID ...
func (*ShowStoreOffer) ID() uint32 {
	return IDShowStoreOffer
}

// Marshal ...
func (pk *ShowStoreOffer) Marshal(io protocol.IO) {
	io.String(&pk.OfferID)
	io.Bool(&pk.ShowAll)
}
