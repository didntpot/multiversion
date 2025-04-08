package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

const (
	NPCRequestActionSetActions = iota
	NPCRequestActionExecuteAction
	NPCRequestActionExecuteClosingCommands
	NPCRequestActionSetName
	NPCRequestActionSetSkin
	NPCRequestActionSetInteractText
	NPCRequestActionExecuteOpeningCommands
)

// NPCRequest is sent by the client when it interacts with an NPC.
// The packet is specifically made for Education Edition, where NPCs are available to use.
type NPCRequest struct {
	EntityRuntimeID uint64
	RequestType     byte
	CommandString   string
	ActionType      byte
	SceneName       string
}

// ID ...
func (*NPCRequest) ID() uint32 {
	return IDNPCRequest
}

// Marshal ...
func (pk *NPCRequest) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Uint8(&pk.RequestType)
	io.String(&pk.CommandString)
	io.Uint8(&pk.ActionType)
	io.String(&pk.SceneName)
}
