package main

import (
	"flag"
	"html/template"
	"log"
)

// The main function initializes a cache, parses command line flags to set configuration values,
// creates an App instance, logs server information, and starts the application.
func main() {
	cache := make(map[string]*template.Template)

	config := Config{
		Version: "1.0.0",
	}

	flag.StringVar(&config.Port, "port", "8080", "HTTP Port")
	flag.StringVar(&config.Env, "env", "dev", "Environment")
	flag.Parse()

	app := App{
		Config: config,
		Cache:  cache,
	}

	log.Printf("Server running on port %s Version: %s Environment: %s", config.Port, config.Version, config.Env)
	log.Fatal(app.Start())
}
