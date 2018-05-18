package cmd

import (
	"github.com/jorgechato/acictl/utils"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{Use: utils.NAME}

	verbose bool
	debug   bool

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize an Acid working directory",
		Long:  "",
		Args:  cobra.NoArgs,
		Run:   initialize,
	}

	configCmd = &cobra.Command{
		Use:   "config [APPNAME]",
		Short: "Configure",
		Long:  "",
		Args:  cobra.NoArgs,
		Run:   config,
	}

	repoPtr string
)

func init() {
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

	configCmd.Flags().StringVarP(
		&repoPtr,
		"repo",
		"r",
		"",
		"Repository where to store the config files",
	)
	configCmd.MarkFlagRequired("repo")
}
