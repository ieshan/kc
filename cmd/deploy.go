package cmd

import (
	"github.com/ieshan/kc/helper"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Returns a list of deployments",
	RunE: func(cmd *cobra.Command, args []string) error {
		kubeCmd := exec.Command("kubectl", "get", "deploy", "-n", namespace)
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err := kubeCmd.Run()
		return helper.ErrorPrintln(err, "Command execution error")
	},
}

func init()  {
	rootCmd.AddCommand(deployCmd)
}
