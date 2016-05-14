package Configuration

import (
	"fmt"

	"gopkg.in/fsnotify.v1"

	"github.com/spf13/viper"
)

var (
	// Port - Port the webserver listens on
	Port string
)

// InitConfig fetch config file
func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("...")
	viper.AddConfigPath("..")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		getSetValues()
	})

	getSetValues()
}

// Parses the config and sets the values
func getSetValues() {
	Port = viper.GetString("port")
}
