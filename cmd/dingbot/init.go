package dingbot

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "init config",
	Long:  `init config file`,
	Run: func(cmd *cobra.Command, args []string) {
		var configFile = args[0]

		err := os.MkdirAll(path.Dir(configFile), 0755)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		viper.SetConfigFile(configFile)

		viper.Set("token", dingTalk.AccessToken)
		viper.Set("secret", dingTalk.Secret)

		err = viper.WriteConfig()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}
