package resources

import (
	"html/template"
	"path"
	"strings"
)

var CompiledTemplates = compileTemplates()

func compileTemplates() *template.Template {
	rootTemplate := template.New(`root_template`)
	for _, assetName := range discoverAllTemplatePaths() {
		subTemplate := rootTemplate.New(path.Base(assetName))
		subTemplate.Parse(string(MustAsset(assetName)))
	}
	return rootTemplate
}

func discoverAllTemplatePaths() []string {
	paths := []string{}
	for _, name := range AssetNames() {
		if strings.Contains(name, "template") {
			paths = append(paths, name)
		}
	}
	return paths
}
