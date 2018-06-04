package cmd

import (
	"fmt"

	"github.com/jorgechato/acictl/utils"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: utils.NAME,
		Long: fmt.Sprintf(
			`%v v%v is a CLI to control the acid project.
This application is a tool to generate the needed files
to quickly create a Jenkins as a service environment in K8S.`,
			utils.NAME,
			utils.VERSION,
		),
		Run: root,
	}

	version bool
	verbose bool
	debug   bool
)

func init() {
	rootCmd.Flags().BoolVarP(
		&version,
		"version",
		"V",
		false,
		"acictl version",
	)

	rootCmd.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"verbose output",
	)

	rootCmd.PersistentFlags().BoolVarP(
		&debug,
		"debug",
		"d",
		false,
		"debug output",
	)

	deployPtr()
	configPtr()
}
