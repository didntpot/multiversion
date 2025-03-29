package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	CommandOutputTypeNone = iota
	CommandOutputTypeLastOutput
	CommandOutputTypeSilent
	CommandOutputTypeAllOutput
	CommandOutputTypeDataSet
)

// CommandOutput is sent by the server to the client to send text as output of a command. Most servers do not
// use this packet and instead simply send Text packets, but there is reason to send it.
// If the origin of a CommandRequest packet is not the player itself, but, for example, a websocket server,
// sending a Text packet will not do what is expected: The message should go to the websocket server, not to
// the client's chat. The CommandOutput packet will make sure the messages are relayed to the correct origin
// of the command request.
type CommandOutput struct {
	CommandOrigin  protocol.CommandOrigin
	OutputType     byte
	SuccessCount   uint32
	OutputMessages []protocol.CommandOutputMessage
	DataSet        string
}

// ID ...
func (*CommandOutput) ID() uint32 {
	return IDCommandOutput
}

// Marshal ...
func (pk *CommandOutput) Marshal(io protocol.IO) {
	l := uint32(len(pk.OutputMessages))

	protocol.CommandOriginData(io, &pk.CommandOrigin)
	io.Uint8(&pk.OutputType)
	io.Varuint32(&pk.SuccessCount)
	io.Varuint32(&l)
	for _, message := range pk.OutputMessages {
		legacyprotocol.CommandMessage(io, (*legacyprotocol.CommandOutputMessage)(&message))
	}
	if pk.OutputType == CommandOutputTypeDataSet {
		io.String(&pk.DataSet)
	}
}
