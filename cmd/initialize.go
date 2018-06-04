package cmd

import (
	"github.com/jorgechato/acictl/generator"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize an Acid working directory",
		Long:  "",
		Args:  cobra.NoArgs,
		Run:   initialize,
	}
)

func initialize(cmd *cobra.Command, args []string) {
	g := generator.New()
	g.DownloadTemplates()
}
