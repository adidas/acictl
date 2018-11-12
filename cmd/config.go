package cmd

import (
	"github.com/adidas/acictl/generator"
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config [APPNAME]",
		Short: "Configure",
		Long:  "",
		Args:  cobra.NoArgs,
		Run:   config,
	}

	repoPtr  string
	tokenPtr string
)

func configPtr() {
	configCmd.Flags().StringVarP(
		&repoPtr,
		"repo",
		"r",
		"",
		"Repository where to store the config files",
	)

	configCmd.Flags().StringVarP(
		&tokenPtr,
		"token",
		"t",
		"",
		"Token with the right to perform github operations",
	)

	configCmd.MarkFlagRequired("token")
}

func config(cmd *cobra.Command, args []string) {
	generator.NewConfig("", tokenPtr)
}
