package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// The Config type represents a configuration object with properties for port, environment, and
// version.
// @property {string} Port - The `Port` property is a string that represents the port number on which
// the application will listen for incoming requests.
// @property {string} Env - The `Env` property in the `Config` struct represents the environment in
// which the application is running. It can be used to differentiate between development, staging, and
// production environments, for example.
// @property {string} Version - The "Version" property in the Config struct represents the version of
// the application or system. It is typically used to track and manage different versions of the
// software.
type Config struct {
	Port    string
	Env     string
	Version string
}

// The App type represents an application with a configuration and a cache of template objects.
// @property {Config} Config - The `Config` property is a variable of type `Config`. It is used to
// store configuration settings for the app.
// @property Cache - The `Cache` property is a map that stores compiled templates. It uses string keys
// to identify the templates and the corresponding values are pointers to the compiled template
// objects. This allows for efficient retrieval and reuse of templates in the application.
type App struct {
	Config Config
	Cache  map[string]*template.Template
}

// The `Start` method is a function defined on the `App` struct. It starts the HTTP server and listens
// for incoming requests on the specified port.
func (a *App) Start() error {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", a.Config.Port),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler: a.Routes(),
	}

	return srv.ListenAndServe()
}
