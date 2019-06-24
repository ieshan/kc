package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var pfCmd = &cobra.Command{
	Use:   "pf",
	Short: "Forward port: example: kc pf {{pod}} {{local_port}}:{{pod_port}}",
	Run: func(cmd *cobra.Command, args []string) {
		kubeCmd := exec.Command("kubectl", "port-forward", args[0], "-n", namespace, args[1])
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err := kubeCmd.Run()
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	},
}

func init()  {
	rootCmd.AddCommand(pfCmd)
}
