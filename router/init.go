package router

import (
	"html/template"
)

var (
	temp *template.Template
	host string
)

func Init(s string) {
	host = s
	temp = template.Must(template.ParseGlob("template/*html"))
}
