package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft"
)

// listenerFunc ...
func listenerFunc(c *server.Config, addr string, protocols []minecraft.Protocol) {
	c.Listeners = []func(conf server.Config) (server.Listener, error){
		func(conf server.Config) (server.Listener, error) {
			cfg := minecraft.ListenConfig{
				MaximumPlayers:         conf.MaxPlayers,
				StatusProvider:         conf.StatusProvider,
				AuthenticationDisabled: conf.AuthDisabled,
				ResourcePacks:          conf.Resources,
				Biomes:                 biomes(),
				TexturePacksRequired:   conf.ResourcesRequired,

				AcceptedProtocols:   protocols,
				AllowUnknownPackets: true,
				AllowInvalidPackets: true,
			}
			if conf.Log.Enabled(context.Background(), slog.LevelDebug) {
				cfg.ErrorLog = conf.Log.With("net origin", "gophertunnel")
			}
			l, err := cfg.Listen("raknet", addr)
			if err != nil {
				return nil, fmt.Errorf("create minecraft listener: %w", err)
			}
			conf.Log.Info("Listener running.", "addr", l.Addr())
			return listener{l}, nil
		},
	}
}

// ashyBiome represents a biome that has any form of ash.
type ashyBiome interface {
	// Ash returns the ash and white ash of the biome.
	Ash() (ash float64, whiteAsh float64)
}

// sporingBiome represents a biome that has blue or red spores.
type sporingBiome interface {
	// Spores returns the blue and red spores of the biome.
	Spores() (blueSpores float64, redSpores float64)
}

// biomes builds a mapping of all biome definitions of the server, ready to be
// set in the biomes field of the server listener.
func biomes() map[string]any {
	definitions := make(map[string]any)
	for _, b := range world.Biomes() {
		definition := map[string]any{
			"name_hash":   b.String(), // Not actually a hash despite the name.
			"temperature": float32(b.Temperature()),
			"downfall":    float32(b.Rainfall()),
			"rain":        b.Rainfall() > 0,
		}
		if a, ok := b.(ashyBiome); ok {
			ash, whiteAsh := a.Ash()
			definition["ash"], definition["white_ash"] = float32(ash), float32(whiteAsh)
		}
		if s, ok := b.(sporingBiome); ok {
			blueSpores, redSpores := s.Spores()
			definition["blue_spores"], definition["red_spores"] = float32(blueSpores), float32(redSpores)
		}
		definitions[b.String()] = definition
	}
	return definitions
}
