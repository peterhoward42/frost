package view

import (
	"fmt"
	"html/template"
)

var GlobalGuiTemplate *template.Template = nil

func InitialiseGlobalGuiTemplate() {
	parsedTemplate, err := template.ParseGlob("static/templates/*.html")
	if err != nil {
		panic(fmt.Sprintf("template parsing failed with %v", err.Error()))
	}
	GlobalGuiTemplate = parsedTemplate
}
