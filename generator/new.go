package generator

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/adidas/acictl/utils"
)

// New creates the $HOME/.acid/templates folder and download all requirements
func New() Generator {
	user, _ := user.Current()
	dot := path.Join(
		user.HomeDir,
		fmt.Sprintf(".%v", utils.PROJECT),
	)

	g := Generator{
		Folders: Folder{
			Dot:       dot,
			Templates: path.Join(dot, "templates"),
		},
		Sources: utils.Sources(),
	}

	os.MkdirAll(g.Folders.Templates, 0777)

	return g
}

// DownloadTemplates download the required templates
func (g *Generator) DownloadTemplates() error {
	// TODO
	return nil
}
