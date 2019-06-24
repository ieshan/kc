package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var svcCmd = &cobra.Command{
	Use:   "svc",
	Short: "Returns a list of svc",
	Run: func(cmd *cobra.Command, args []string) {
		kubeCmd := exec.Command("kubectl", "get", "svc", "-n", namespace)
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err := kubeCmd.Run()
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	},
}

func init()  {
	rootCmd.AddCommand(svcCmd)
}
