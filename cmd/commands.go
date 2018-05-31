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
			"%v v%v is a CLI to control the acid project.\nThis environment is a tool to generate the needed files\nto quickly create a Jenkins as a service application in K8S.",
			utils.NAME,
			utils.VERSION,
		),
		Run: root,
	}

	version bool
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
	rootCmd.Flags().BoolVarP(
		&verbose,
		"version",
		"V",
		true,
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

	configCmd.Flags().StringVarP(
		&repoPtr,
		"repo",
		"r",
		"",
		"Repository where to store the config files",
	)
	configCmd.MarkFlagRequired("repo")
}
