package protocol

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// Attribute is an entity attribute, that holds specific data such as the health of the entity. Each attribute
// holds a default value, maximum and minimum value, name and its current value.
type Attribute struct {
	protocol.Attribute
}

// Marshal ...
func (x *Attribute) Marshal(io protocol.IO) {
	io.Float32(&x.Min)
	io.Float32(&x.Max)
	io.Float32(&x.Value)
	io.Float32(&x.Default)
	io.String(&x.Name)
}
