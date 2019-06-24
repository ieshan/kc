package cmd

import (
	"github.com/ieshan/kc/helper"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Returns a list of pods",
	RunE: func(cmd *cobra.Command, args []string) error {
		kubeCmd := exec.Command("kubectl", "get", "pods", "-n", namespace)
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err := kubeCmd.Run()
		return helper.ErrorPrintln(err, "Command execution error")
	},
}

func init()  {
	rootCmd.AddCommand(podsCmd)
}
