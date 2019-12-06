package main

import (
	"flag"
	"fmt"
	"html"

	"log"
	"net/http"
	"time"

	"github.com/adamisrael/gedcom-api/internal/router"
	. "github.com/adamisrael/gedcom-api/internal/utils"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	return
}

func main() {

	verbose := flag.Bool("verbose", false, "verbosity")
	flag.Parse()

	configuration := GetConfig()

	// viper.SetConfigName("gedcom-api")
	// viper.AddConfigPath(".")
	// viper.AddConfigPath("../../configs")
	// var configuration config.Configuration
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file, %s", err)
	// }
	// err := viper.Unmarshal(&configuration)
	// if err != nil {
	// 	log.Fatalf("unable to decode into struct, %v", err)
	// }

	r := router.GetRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", configuration.Server.Hostname, configuration.Server.Port),
		WriteTimeout: configuration.Server.WriteTimeout * time.Second,
		ReadTimeout:  configuration.Server.ReadTimeout * time.Second,
	}

	if *verbose {
		fmt.Printf("Listening on %s:%d\n", configuration.Server.Hostname, configuration.Server.Port)
		fmt.Printf("Read Timeout %ds\n", configuration.Server.ReadTimeout)
		fmt.Printf("Write Timeout %ds\n", configuration.Server.WriteTimeout)
		fmt.Printf("GEDCOM path: %s\n", configuration.Gedcom.Path)
		fmt.Printf("Max Upload Size: %dM\n", configuration.Gedcom.MaxFileSize)
	}

	log.Fatal(srv.ListenAndServe())

}
