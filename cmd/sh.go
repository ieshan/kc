package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var shCmd = &cobra.Command{
	Use:   "sh",
	Short: "Access pod shell",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kubeCmd := exec.Command("kubectl", "exec", "-it", args[0], "-n", namespace, "--", "sh")
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		kubeCmd.Stdin = os.Stdin
		err := kubeCmd.Run()
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	},
}

func init()  {
	rootCmd.AddCommand(shCmd)
}
