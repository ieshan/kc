package namespace

import (
	"github.com/ieshan/kc/helper"
	"github.com/spf13/cobra"
	"sort"
)

var nsSaveCmd = &cobra.Command{
	Use:   "add",
	Short: "Add namespace",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nsList := helper.GetStringSlice("NamespaceList")
		if _, inList := helper.Contains(args[0], nsList); inList {
			return
		}
		nsList = append(nsList, args[0])
		sort.Strings(nsList)
		helper.Set("NamespaceList", nsList)
		helper.WriteConfig()
	},
}

func init()  {
	NsCmd.AddCommand(nsSaveCmd)
}

