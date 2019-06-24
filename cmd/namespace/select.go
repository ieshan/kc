package namespace

import (
	"github.com/ieshan/kc/helper"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var nsListCmd = &cobra.Command{
	Use:   "select",
	Short: "Select namespace",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Select Default Namespace (" + helper.GetString("DefaultNamespace") + "):",
			Items: helper.GetStringSlice("NamespaceList"),
		}
		_, result, err := prompt.Run()
		helper.ErrorFatal(err, "Prompt failed.\n")
		helper.Set("DefaultNamespace", result)
		helper.WriteConfig()
	},
}

func init()  {
	NsCmd.AddCommand(nsListCmd)
}

