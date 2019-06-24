package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Returns a list of pods",
	Run: func(cmd *cobra.Command, args []string) {
		kubeCmd := exec.Command("kubectl", "get", "pods", "-n", namespace)
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err := kubeCmd.Run()
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	},
}

func init()  {
	rootCmd.AddCommand(podsCmd)
}
