package utils

import "fmt"

var VERSION string = "1.0"

const (
	NAME    string = "acictl"
	PROJECT string = "acid"
	// COMPANY string = "adidas"
	COMPANY string = "jorgechato"
)

func Templates() []string {
	return []string{
		"generator/init/templates/config.xml.tmpl",
	}
}

func Repos() []string {
	return []string{
		"master-image-base",
		"slave-image-base",
		"master-image-deploy",
		"master-config-changeset",
		"master-config",
	}
}

func Sources() []string {
	sources := []string{}

	for _, val := range Templates() {
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

func Error() errors {
	return errors{
		Client: "ERROR: client failed [%v]",
	}
}

func Msg() messages {
	return messages{
		Init:      "Initialize an Acid working directory",
		Download:  "Download and install modules for the configuration",
		Canonical: "Rewrites config files to canonical format",
	}
}

func Repo() repo {
	return repo{
		Instance: "config-%v",
		Base:     "config-base",
	}
}
