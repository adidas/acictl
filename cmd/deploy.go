package cmd

import (
	"fmt"

	"github.com/jorgechato/acictl/k8s"
	"github.com/spf13/cobra"
)

var (
	deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "Deploy",
		Long:  "",
		Args:  cobra.NoArgs,
		Run:   deploy,
	}

	kubeconfigPtr string
)

func deployPtr() {
	deployCmd.Flags().StringVarP(
		&kubeconfigPtr,
		"kubeconfig",
		"",
		"",
		"Path to the kubeconfig file to use for CLI requests",
	)
	deployCmd.MarkFlagRequired("kubeconfig")
}

func deploy(cmd *cobra.Command, args []string) {
	c := k8s.NewClient(kubeconfigPtr)
	fmt.Println(c)
}
