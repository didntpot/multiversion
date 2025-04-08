package protocol

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// Command holds the data that a command requires to be shown to a player client-side. The command is shown in
// the /help command and auto-completed using this data.
type Command struct {
	Name            string
	Description     string
	Flags           uint16
	PermissionLevel byte
	AliasesOffset   uint32
	Overloads       []CommandOverload
}

// Marshal ...
func (x *Command) Marshal(io protocol.IO) {
	io.String(&x.Name)
	io.String(&x.Description)
	io.Uint16(&x.Flags)
	io.Uint8(&x.PermissionLevel)
	io.Uint32(&x.AliasesOffset)
	protocol.Slice(io, &x.Overloads)
}

// CommandOverload represents an overload of a command. This overload can be compared to function overloading
// in languages such as java. It represents a single usage of the command. A command may have multiple
// different overloads, which are handled differently.
type CommandOverload struct {
	Parameters []protocol.CommandParameter
}

// Marshal ...
func (x *CommandOverload) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.Parameters)
}

// CommandOutputMessage represents a message sent by a command that holds the output of one of the commands
// executed.
type CommandOutputMessage struct {
	Success    bool
	Message    string
	Parameters []string
}

// CommandMessage ...
func CommandMessage(io protocol.IO, x *CommandOutputMessage) {
	l := uint32(len(x.Parameters))

	io.Bool(&x.Success)
	io.String(&x.Message)
	io.Varuint32(&l)
	for _, param := range x.Parameters {
		io.String(&param)
	}
}
