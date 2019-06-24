package cmd

import (
	"github.com/spf13/cobra"
)

var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Work with kubernates namespace",
	Run: func(cmd *cobra.Command, args []string) {},
}

func init()  {
	rootCmd.AddCommand(nsCmd)
}
