package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	TextTypeRaw = iota
	TextTypeChat
	TextTypeTranslation
	TextTypePopup
	TextTypeJukeboxPopup
	TextTypeTip
	TextTypeSystem
	TextTypeWhisper
	TextTypeAnnouncement
	TextTypeObjectWhisper
	TextTypeObject
	TextTypeObjectAnnouncement
)

// Text is sent by the client to the server to send chat messages, and by the server to the client to forward
// or send messages, which may be chat, popups, tips etc.
type Text struct {
	TextType         byte
	NeedsTranslation bool
	SourceName       string
	Message          string
	Parameters       []string
	XUID             string
	PlatformChatID   string
}

// ID ...
func (*Text) ID() uint32 {
	return IDText
}

// Marshal ...
func (pk *Text) Marshal(io protocol.IO) {
	io.Uint8(&pk.TextType)
	io.Bool(&pk.NeedsTranslation)
	switch pk.TextType {
	case TextTypeChat, TextTypeWhisper, TextTypeAnnouncement:
		io.String(&pk.SourceName)
		io.String(&pk.Message)
	case TextTypeRaw, TextTypeTip, TextTypeSystem, TextTypeObject, TextTypeObjectWhisper, TextTypeObjectAnnouncement:
		io.String(&pk.Message)
	case TextTypeTranslation, TextTypePopup, TextTypeJukeboxPopup:
		io.String(&pk.Message)
		protocol.FuncSlice(io, &pk.Parameters, io.String)
	}
	io.String(&pk.XUID)
	io.String(&pk.PlatformChatID)
}
