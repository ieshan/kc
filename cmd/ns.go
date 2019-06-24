package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Save namespace",
	Run: func(cmd *cobra.Command, args []string) {
		err := ioutil.WriteFile(".kubectl_namespace", []byte(args[0]), 0666)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	},
}

func init()  {
	rootCmd.AddCommand(nsCmd)
}
