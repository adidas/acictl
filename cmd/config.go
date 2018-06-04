package cmd

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/jorgechato/acictl/utils/httpclient"
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

	repoPtr string
)

func configPtr() {
	configCmd.Flags().StringVarP(
		&repoPtr,
		"repo",
		"r",
		"",
		"Repository where to store the config files",
	)
	configCmd.MarkFlagRequired("repo")
}

func config(cmd *cobra.Command, args []string) {
	c := httpclient.NewClient(
		"https://api.github.com",
		"GET",
		"application/json",
		"",
		"",
	)

	c.Run(
		bytes.Buffer{},
		func(res *http.Response, b string) {

		},
		verbose,
		debug,
	)

	fmt.Print("With love by adidas.")
}
