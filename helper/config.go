package helper

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

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
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}

func GetString(key  string) string {
	return viper.GetString(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func WriteConfig() {
	homeDir, err := os.UserHomeDir()
	ErrorFatal(err, "Could not find $HOME dir")
	ErrorFatal(viper.WriteConfigAs(strings.TrimRight(homeDir, "/") + "/.kc_config.json"), "Couldn't write config file.\n")
}
