package cmd

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/jorgechato/acictl/utils/httpclient"
	"github.com/spf13/cobra"
)

func config(cmd *cobra.Command, args []string) {
	c := httpclient.NewClient(
		"https://api.github.com",
		"GET",
		"application/json",
		"",
		"",
	)

	err := c.Run(
		bytes.Buffer{},
		func(res *http.Response, b string) {
		},
		verbose,
		debug,
	)

	fmt.Print(err)
	fmt.Print("With love by adidas.")
}
