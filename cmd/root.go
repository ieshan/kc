package cmd

import (
	"github.com/ieshan/kc/helper"
	ns "github.com/ieshan/kc/cmd/namespace"
	"github.com/spf13/cobra"
)

var namespace string

var rootCmd = &cobra.Command{
	Use:   "kc",
	Short: "Basic kubernetes commands",
	Long: ``,
}

func init()  {
	helper.Setup()
	defaultNamespace := helper.GetString("DefaultNamespace")
	rootCmd.AddCommand(ns.NsCmd)
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", defaultNamespace, "namespace")
}

func Execute() {
	helper.ErrorFatal(rootCmd.Execute(), "Failed command execution")
}
