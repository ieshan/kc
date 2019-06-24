package namespace

import (
	"github.com/ieshan/kc/helper"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var nsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete namespace",
	Run: func(cmd *cobra.Command, args []string) {
		nsList := helper.GetStringSlice("NamespaceList")
		prompt := promptui.Select{
			Label: "Select namespace to delete:",
			Items: nsList,
		}
		_, result, err := prompt.Run()
		helper.ErrorFatal(err, "Prompt failed.\n")
		if result == "default" {
			return
		}
		pos, inList := helper.Contains(result, nsList)
		if !inList {
			return
		}
		copy(nsList[pos:], nsList[pos+1:]) // Shift a[i+1:] left one index.
		nsList[len(nsList) - 1] = ""     // Erase last element (write zero value).
		nsList = nsList[:len(nsList)-1]

		helper.Set("NamespaceList", nsList)
		if helper.GetString("DefaultNamespace") == result {
			helper.Set("DefaultNamespace", "default")
		}
		helper.WriteConfig()
	},
}

func init()  {
	NsCmd.AddCommand(nsDeleteCmd)
}
