package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var dpCmd = &cobra.Command{
	Use:   "dp",
	Short: "Delete pod",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kubeCmd := exec.Command("kubectl", "delete", "pods", "-n", namespace, args[0])
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err := kubeCmd.Run()
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	},
}

func init()  {
	rootCmd.AddCommand(dpCmd)
}
