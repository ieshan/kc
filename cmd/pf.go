package cmd

import (
	"github.com/ieshan/kc/helper"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var pfCmd = &cobra.Command{
	Use:   "pf",
	Short: "Forward port: example: kc pf {{pod}} {{local_port}}:{{pod_port}}",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		kubeCmd := exec.Command("kubectl", "port-forward", args[0], "-n", namespace, args[1])
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err := kubeCmd.Run()
		return helper.ErrorPrintln(err, "Command execution error")
	},
}

func init()  {
	rootCmd.AddCommand(pfCmd)
}
