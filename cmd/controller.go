package cmd

import (
	"fmt"

	"github.com/jorgechato/acictl/generator"
	"github.com/spf13/cobra"
)

func initialize(cmd *cobra.Command, args []string) {
	g := generator.New()
	g.DownloadTemplates()
}

func config(cmd *cobra.Command, args []string) {
	fmt.Print("With love by adidas.")
}
