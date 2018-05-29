package generator

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"testing"

	"github.com/jorgechato/acictl/utils"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	user, _ := user.Current()
	dot := path.Join(
		user.HomeDir,
		fmt.Sprintf(".%v", utils.PROJECT),
	)

	mockG := Generator{
		Folders: Folder{
			Dot:       dot,
			Templates: path.Join(dot, "templates"),
		},
		Sources: utils.Sources(),
	}

	g := New()

	assert.Equal(mockG, g, "Models should be equals")

	fi, err := os.Lstat(g.Folders.Templates)
	mode := fi.Mode()

	assert.Nil(err, "Error should be nil")
	assert.True(mode.IsDir(), "Should be a directory")
	assert.Equal("drwxr-xr-x", mode.String(), "Permissions should be 0777")
}

func TestDownloadTemplates(t *testing.T) {
	g := New()
	err := g.DownloadTemplates()
	assert.Nil(t, err, "Error should be nil")
}
