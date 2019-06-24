package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var namespace string

var rootCmd = &cobra.Command{
	Use:   "kc",
	Short: "Basic kubernetes commands",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {},
}

func init()  {
	viper.SetConfigType("json")
	viper.SetConfigName(".kc_config")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")

	viper.SetDefault("DefaultNamespace", "default")
	viper.SetDefault("NamespaceList", []string{"default"})

	err := viper.ReadInConfig()
	if err != nil {
		WriteConfig()
	}
	defaultNamespace := viper.GetString("DefaultNamespace")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", defaultNamespace, "namespace")
}

func Contains(item string, itemList []string) (int, bool)  {
	for i, n := range itemList {
		if item == n {
			return i, true
		}
	}
	return -1, false
}

func Execute() {
	ErrorFatal(rootCmd.Execute(), "Failed command execution")
}

func ErrorFatal(err error, message string) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func ErrorPrintln(err error, message string)  {
	if err != nil {
		log.Println(message, err)
	}
}

func WriteConfig() {
	homeDir, err := os.UserHomeDir()
	ErrorFatal(err, "Could not find $HOME dir")
	ErrorFatal(viper.WriteConfigAs(strings.TrimRight(homeDir, "/") + "/.kc_config.json"), "Couldn't write config file.\n")
}

