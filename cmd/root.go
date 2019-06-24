package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var namespace string

var rootCmd = &cobra.Command{
	Use:   "kc",
	Short: "Basic kubernetes commands",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {},
}

func init()  {
	var defaultNamespace string
	namespaceData, err := ioutil.ReadFile(".kubectl_namespace")
	if err != nil {
		defaultNamespace = "default"
	} else {
		defaultNamespace = string(namespaceData)
	}
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", defaultNamespace, "namespace")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
