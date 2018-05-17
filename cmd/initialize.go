package cmd

import (
	"github.com/jorgechato/acictl/generator"
	"github.com/spf13/cobra"
)

func initialize(cmd *cobra.Command, args []string) {
	g := generator.New()
	g.DownloadTemplates()
}
