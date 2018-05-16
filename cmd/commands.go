package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jorgechato/acictl/utils"
)

var (
	rootCmd = &cobra.Command{Use: utils.NAME}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize an Acid working directory",
		Long:  "",
		Args:  cobra.ExactArgs(1),
		Run:   initialize,
	}

	configCmd = &cobra.Command{
		Use:   "config [APPNAME]",
		Short: "Configure",
		Long:  "",
		Args:  cobra.ExactArgs(1),
		Run:   config,
	}

	verbose bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"verbose output",
	)
}
