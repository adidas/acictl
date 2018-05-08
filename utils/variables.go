package utils

import "fmt"

var (
	NAME    = "acictl"
	PROJECT = "acid"
	// COMPANY = "adidas"
	COMPANY = "jorgechato"
	SOURCES = []string{
		"https://github.com/%v/%v-master-image-base",
		"https://github.com/%v/%v-slave-image-base",
		"https://github.com/%v/%v-master-image-deploy",
		"https://github.com/%v/%v-master-config-changeset",
		"https://github.com/%v/%v-master-config",
	}
)

func init() {
	for i, source := range SOURCES {
		SOURCES[i] = fmt.Sprintf(
			source,
			COMPANY,
			PROJECT,
		)
	}
}
