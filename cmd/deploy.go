package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Returns a list of deployments",
	Run: func(cmd *cobra.Command, args []string) {
		kubeCmd := exec.Command("kubectl", "get", "deploy", "-n", namespace)
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err := kubeCmd.Run()
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	},
}

func init()  {
	rootCmd.AddCommand(deployCmd)
}
