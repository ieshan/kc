package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var logCmd = &cobra.Command{
	Use:   "logs",
	Short: "Pod log",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kubeCmd := exec.Command("kubectl", "logs", "-f", "-n", namespace, args[0])
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err := kubeCmd.Run()
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	},
}

func init()  {
	rootCmd.AddCommand(logCmd)
}
