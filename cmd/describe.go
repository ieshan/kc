package cmd

import (
	"github.com/ieshan/kc/helper"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe pod",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		kubeCmd := exec.Command("kubectl", "describe", "pods", "-n", namespace, args[0])
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err := kubeCmd.Run()
		return helper.ErrorPrintln(err, "Command execution error")
	},
}

func init()  {
	rootCmd.AddCommand(describeCmd)
}
