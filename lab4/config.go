package main

import "html/template"

type Config struct {
	Port    string
	Env     string
	Version string
}

type App struct {
	Config Config
	Cache  map[string]*template.Template
}
