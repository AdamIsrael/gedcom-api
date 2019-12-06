package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/adamisrael/gedcom-api/internal/config"
	"github.com/spf13/viper"
)

// GetConfig returns the application configuration
func GetConfig() config.Configuration {
	viper.SetConfigName("gedcom-api")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../configs")
	var configuration config.Configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if CheckErr(err) {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	// TODO: Restart the http server if the config changes
	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name)
	// })
	return configuration
}

// FileExists checks to see if the filename exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// CheckErr evaluates the error code. Returns a boolean indicating success
// or failure, and prints the error in the later case.
func CheckErr(err error) bool {
	// TODO: Write a checkErr function
	if err != nil {
		// handle error
		fmt.Println(err)
		return true
	}
	return false
}
