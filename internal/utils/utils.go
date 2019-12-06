package utils

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/adamisrael/gedcom-api/internal/config"
	"github.com/google/uuid"
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

// GetGedcomFilename returns the path to use to store an uploaded GEDCOM file
func GetGedcomFilename() (string, string) {
	id, err := uuid.NewUUID()
	CheckErr(err)

	config := GetConfig()
	filename := path.Join(config.Gedcom.Path, id.String()+".ged")
	if FileExists(filename) {
		return GetGedcomFilename()
	}
	return id.String(), filename
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
