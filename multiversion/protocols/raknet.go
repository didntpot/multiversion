package protocols

import (
	"log/slog"
	"net"

	"github.com/sandertv/go-raknet"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// MultiRakNet ...
type MultiRakNet struct {
	minecraft.RakNet
}

// legacyVersions ...
var legacyVersions = []byte{10}

// Listen ...
func (MultiRakNet) Listen(address string) (minecraft.NetworkListener, error) {
	return raknet.ListenConfig{
		ProtocolVersions: legacyVersions,
	}.Listen(address)
}

// Compression ...
func (MultiRakNet) Compression(net.Conn) packet.Compression {
	return packet.FlateCompression
}

// init registers the MultiRakNet network. It overrides the existing minecraft.RakNet network.
func init() {
	minecraft.RegisterNetwork("raknet", func(*slog.Logger) minecraft.Network { return MultiRakNet{} })
}
