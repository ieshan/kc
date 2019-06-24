package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var nsListCmd = &cobra.Command{
	Use:   "select",
	Short: "Select namespace",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Select Default Namespace (" + viper.GetString("DefaultNamespace") + "):",
			Items: viper.GetStringSlice("NamespaceList"),
		}
		_, result, err := prompt.Run()
		ErrorFatal(err, "Prompt failed.\n")
		viper.Set("DefaultNamespace", result)
		WriteConfig()
	},
}

func init()  {
	nsCmd.AddCommand(nsListCmd)
}
