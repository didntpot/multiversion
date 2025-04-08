package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

const (
	NPCDialogueActionOpen int32 = iota
	NPCDialogueActionClose
)

// NPCDialogue is a packet that allows the client to display dialog boxes for interacting with NPCs.
type NPCDialogue struct {
	ActorUniqueID uint64
	ActionType    int32
	Dialogue      string
	SceneName     string
	NPCName       string
	ActionJSON    string
}

// ID ...
func (*NPCDialogue) ID() uint32 {
	return IDNPCDialogue
}

// Marshal ...
func (pk *NPCDialogue) Marshal(io protocol.IO) {
	io.Uint64(&pk.ActorUniqueID)
	io.Varint32(&pk.ActionType)
	io.String(&pk.Dialogue)
	io.String(&pk.SceneName)
	io.String(&pk.NPCName)
	io.String(&pk.ActionJSON)
}
