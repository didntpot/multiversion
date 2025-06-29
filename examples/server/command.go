package main

import (
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

func init() {
	cmd.Register(cmd.New("gamemode", "", nil, GameMode{}))
}

// GameMode ...
type GameMode struct {
	GameMode gameMode `cmd:"gamemode"`
}

// Run ...
func (com GameMode) Run(src cmd.Source, _ *cmd.Output, _ *world.Tx) {
	p := src.(*player.Player)

	var name string
	var mode world.GameMode
	switch strings.ToLower(string(com.GameMode)) {
	case "survival", "0", "s":
		name, mode = "survival", world.GameModeSurvival
	case "creative", "1", "c":
		name, mode = "creative", world.GameModeCreative
	case "adventure", "2", "a":
		name, mode = "adventure", world.GameModeAdventure
	case "spectator", "3", "sp":
		name, mode = "spectator", world.GameModeSpectator
	}
	p.SetGameMode(mode)
	p.Messagef(text.Colourf("<white>You've updated your gamemode to <green>%s</green>.</white>", name))
}

// gameMode ...
type gameMode string

// Type ...
func (gameMode) Type() string {
	return "GameMode"
}

// Options ...
func (gameMode) Options(cmd.Source) []string {
	return []string{
		"survival", "0", "s",
		"creative", "1", "c",
		"adventure", "2", "a",
		"spectator", "3", "sp",
	}
}
