package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var nsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete namespace",
	Run: func(cmd *cobra.Command, args []string) {
		nsList := viper.GetStringSlice("NamespaceList")
		prompt := promptui.Select{
			Label: "Select namespace to delete:",
			Items: nsList,
		}
		_, result, err := prompt.Run()
		ErrorFatal(err, "Prompt failed.\n")
		if result == "default" {
			return
		}
		pos, inList := Contains(result, nsList)
		if !inList {
			return
		}
		copy(nsList[pos:], nsList[pos+1:]) // Shift a[i+1:] left one index.
		nsList[len(nsList) - 1] = ""     // Erase last element (write zero value).
		nsList = nsList[:len(nsList)-1]

		viper.Set("NamespaceList", nsList)
		if viper.GetString("DefaultNamespace") == result {
			viper.Set("DefaultNamespace", "default")
		}
		WriteConfig()
	},
}

func init()  {
	nsCmd.AddCommand(nsDeleteCmd)
}
