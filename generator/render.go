package generator

import (
	// "encoding/xml"

	"github.com/adidas/acictl/generator/github"
)

// NewConfig Generate a new config.xml with the base configuration.
func NewConfig(path, token string) {
	g := github.New(token)
	g.CreateRepo("acictl", "foo")
}
