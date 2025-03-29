package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// EducationSettings is a packet sent by the server to update Minecraft: Education Edition related settings.
// It is unused by the normal base game.
type EducationSettings struct {
	CodeBuilderDefaultURI string
	CodeBuilderTitle      string
	CanResizeCodeBuilder  bool
	DisableLegacyTitleBar bool
	PostProcessFilter     string
	ScreenshotBorderPath  string
	AgentCapabilities     *bool
	OverrideURI           string
	HasQuiz               bool
	ExternalLinkSettings  *protocol.EducationExternalLinkSettings
}

// ID ...
func (*EducationSettings) ID() uint32 {
	return IDEducationSettings
}

// Marshal ...
func (pk *EducationSettings) Marshal(io protocol.IO) {
	hasOverrideURI := pk.OverrideURI != ""
	hasAgentCapabilities := pk.AgentCapabilities != nil
	hasExternalLinkSettings := pk.ExternalLinkSettings != nil

	io.String(&pk.CodeBuilderDefaultURI)
	io.String(&pk.CodeBuilderTitle)
	io.Bool(&pk.CanResizeCodeBuilder)
	io.Bool(&pk.DisableLegacyTitleBar)
	io.String(&pk.PostProcessFilter)
	io.String(&pk.ScreenshotBorderPath)

	io.Bool(&hasAgentCapabilities)
	if hasAgentCapabilities {
		io.Bool(pk.AgentCapabilities)
	}

	io.Bool(&hasOverrideURI)
	if hasOverrideURI {
		io.String(&pk.OverrideURI)
	}

	io.Bool(&pk.HasQuiz)

	io.Bool(&hasExternalLinkSettings)
	if hasExternalLinkSettings {
		io.Bool(&hasExternalLinkSettings)
		io.String(&pk.ExternalLinkSettings.URL)
		io.String(&pk.ExternalLinkSettings.DisplayName)
	}
}
