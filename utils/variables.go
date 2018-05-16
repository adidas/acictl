package utils

var (
	NAME    = "acictl"
	PROJECT = "acid"
	// COMPANY = "adidas"
	COMPANY   = "jorgechato"
	TEMPLATES = []string{
		"master-image-base",
		"slave-image-base",
		"master-image-deploy",
		"master-config-changeset",
		"master-config",
	}
	SOURCES = []string{}

	MSG = map[string]string{
		"init":      "Initialize an Acid working directory",
		"download":  "Download and install modules for the configuration",
		"canonical": "Rewrites config files to canonical format",
	}
)
