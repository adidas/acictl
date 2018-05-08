package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jorgechato/acictl/utils"
)

var (
	rootCmd = &cobra.Command{Use: utils.NAME}

	configCmd = &cobra.Command{
		Use:   "config [APPNAME]",
		Short: "",
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
