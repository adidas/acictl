package utils

import "fmt"

var (
	NAME    = "acictl"
	PROJECT = "acid"
	// COMPANY = "adidas"
	COMPANY = "jorgechato"
	REPOS   = []string{
		"master-image-base",
		"slave-image-base",
		"master-image-deploy",
		"master-config-changeset",
		"master-config",
	}
	TEMPLATES = []string{
		"generator/init/templates/config.xml.tmpl",
	}

	MSG = map[string]string{
		"init":      "Initialize an Acid working directory",
		"download":  "Download and install modules for the configuration",
		"canonical": "Rewrites config files to canonical format",
	}
	ERROR = map[string]string{
		"client": "ERROR: client failed [%v]",
	}
)

func Sources() []string {
	sources := []string{}

	for _, val := range TEMPLATES {
		sources = append(
			sources,
			fmt.Sprintf(
				"https://github.com/downloads/%v/%v/%v",
				COMPANY,
				NAME,
				val,
			),
		)
	}

	return sources
}
