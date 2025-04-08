package main

import (
	"log/slog"

	"github.com/cooldogedev/spectrum"
	"github.com/cooldogedev/spectrum/server"
	"github.com/cooldogedev/spectrum/util"
	v486 "github.com/didntpot/multiversion/multiversion/protocols/v486"
	"github.com/didntpot/multiversion/multiversion/protocols/v486/protocol/packet"
	"github.com/sandertv/gophertunnel/minecraft"
)

import _ "github.com/didntpot/multiversion/multiversion/protocols"

// clientDecode ...
var clientDecode = []uint32{ // TODO: This is the bare minimum to get the client to spawn in
	packet.IDTickSync,
	packet.IDPlayerAuthInput,
	packet.IDText,
	packet.IDRequestChunkRadius,
}

// NOTE: Spectrum requires you enable packet decoding server sided as well
// you can do this via spectrum-pm (for PocketMine), or spectrum-df (for Dragonfly)
// spectrum-pm:
// \cooldogedev\Spectrum\Spectrum->registerPacketDecode(23, true)
// spectrum-df:
// util.RegisterPacketDecode(packet.IDTickSync, true)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	log := slog.Default()

	opts := util.DefaultOpts()
	opts.SyncProtocol = false
	opts.ClientDecode = clientDecode
	proxy := spectrum.NewSpectrum(server.NewStaticDiscovery("127.0.0.1:19133", ""), log, opts, nil)
	if err := proxy.Listen(minecraft.ListenConfig{
		StatusProvider: util.NewStatusProvider("Spectrum Proxy", "Spectrum"),

		AcceptedProtocols: []minecraft.Protocol{
			v486.New(false),
		},
		AllowUnknownPackets: true,
		AllowInvalidPackets: true,
	}); err != nil {
		return
	}

	for {
		_, _ = proxy.Accept()
	}
}
