package cmd

import (
	"github.com/ieshan/kc/helper"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)


var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Select context",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := helper.GetKubeConfig()
		helper.ErrorFatal(err, "Error reading kubenetes config")
		contexts := []string{}
		for _, v := range cfg.Contexts {
			contexts = append(contexts, v.Name)
		}
		prompt := promptui.Select{
			Label: "Select Default Context (" + cfg.CurrentContext + "):",
			Items: contexts,
		}
		_, result, err := prompt.Run()
		helper.ErrorFatal(err, "Prompt failed.\n")
		kubeCmd := exec.Command("kubectl", "config", "use-context", result)
		kubeCmd.Stdout = os.Stdout
		kubeCmd.Stderr = os.Stderr
		err = kubeCmd.Run()
		return helper.ErrorPrintln(err, "Command execution error")
	},
}

func init()  {
	rootCmd.AddCommand(contextCmd)
}

