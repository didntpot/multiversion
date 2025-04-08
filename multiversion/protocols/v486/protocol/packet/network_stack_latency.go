package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// NetworkStackLatency is sent by the server (and the client, on development builds) to measure the latency
// over the entire Minecraft stack, rather than the RakNet latency. It has other usages too, such as the
// ability to be used as some kind of acknowledgement packet, to know when the client has received a certain
// other packet.
type NetworkStackLatency struct {
	Timestamp     int64
	NeedsResponse bool
}

// ID ...
func (*NetworkStackLatency) ID() uint32 {
	return IDNetworkStackLatency
}

// Marshal ...
func (pk *NetworkStackLatency) Marshal(io protocol.IO) {
	io.Int64(&pk.Timestamp)
	io.Bool(&pk.NeedsResponse)
}
