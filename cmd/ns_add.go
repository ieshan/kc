package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sort"
)

var nsSaveCmd = &cobra.Command{
	Use:   "add",
	Short: "Add namespace",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nsList := viper.GetStringSlice("NamespaceList")
		if _, inList := Contains(args[0], nsList); inList {
			return
		}
		nsList = append(nsList, args[0])
		sort.Strings(nsList)
		viper.Set("NamespaceList", nsList)
		WriteConfig()
	},
}

func init()  {
	nsCmd.AddCommand(nsSaveCmd)
}
