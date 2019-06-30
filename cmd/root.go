package cmd

import (
	ns "github.com/ieshan/kc/cmd/namespace"
	"github.com/ieshan/kc/helper"
	"github.com/spf13/cobra"
)

var namespace string

var rootCmd = &cobra.Command{
	Use:   "kc",
	Short: "Basic kubernetes commands",
	Long: ``,
}

func Execute() {
	defaultNamespace := helper.GetString("DefaultNamespace")
	rootCmd.AddCommand(ns.NsCmd)
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", defaultNamespace, "namespace")
	helper.ErrorFatal(rootCmd.Execute(), "Failed command execution")
}
