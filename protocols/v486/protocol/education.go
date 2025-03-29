package protocol

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// EducationSharedResourceURI is an education edition feature that is used for transmitting
// education resource settings to clients. It contains a button name and a link URL.
type EducationSharedResourceURI struct {
	ButtonName string
	LinkURI    string
}

// EducationResourceURI ...
func EducationResourceURI(io protocol.IO, x *EducationSharedResourceURI) {
	io.String(&x.ButtonName)
	io.String(&x.LinkURI)
}

// EducationExternalLinkSettings ...
type EducationExternalLinkSettings struct {
	URL         string
	DisplayName string
}
