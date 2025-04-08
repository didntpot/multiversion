package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	BookActionReplacePage = iota
	BookActionAddPage
	BookActionDeletePage
	BookActionSwapPages
	BookActionSign
)

// BookEdit is sent by the client when it edits a book. It is sent each time a modification was made and the
// player stops its typing 'session', rather than simply after closing the book.
type BookEdit struct {
	ActionType          byte
	InventorySlot       byte
	PageNumber          byte
	SecondaryPageNumber byte
	Text                string
	PhotoName           string
	Title               string
	Author              string
	XUID                string
}

// ID ...
func (*BookEdit) ID() uint32 {
	return IDBookEdit
}

// Marshal ...
func (pk *BookEdit) Marshal(io protocol.IO) {
	io.Uint8(&pk.ActionType)
	io.Uint8(&pk.InventorySlot)
	switch pk.ActionType {
	case BookActionReplacePage, BookActionAddPage:
		io.Uint8(&pk.PageNumber)
		io.String(&pk.Text)
		io.String(&pk.PhotoName)
	case BookActionDeletePage:
		io.Uint8(&pk.PageNumber)
	case BookActionSwapPages:
		io.Uint8(&pk.PageNumber)
		io.Uint8(&pk.SecondaryPageNumber)
	case BookActionSign:
		io.String(&pk.Title)
		io.String(&pk.Author)
		io.String(&pk.XUID)
	default:
		io.UnknownEnumOption(pk.ActionType, "book edit action type")
	}
}
