package cmd

import (
	"github.com/ieshan/kc/helper"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var shCmd = &cobra.Command{
	Use:   "sh",
	Short: "Access pod shell",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		kubeCmd := exec.Command("kubectl", "exec", "-it", args[0], "-n", namespace, "--", "sh")
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		kubeCmd.Stdin = os.Stdin
		err := kubeCmd.Run()
		return helper.ErrorPrintln(err, "Command execution error")
	},
}

func init()  {
	rootCmd.AddCommand(shCmd)
}
