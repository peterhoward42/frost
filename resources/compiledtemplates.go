package resources

import (
	"html/template"
	"path"
	"strings"
)

// The CompiledTemplates package-global variable provides access to the html template for
// rendering the gui, having extracted the template files from the behind-the-scenes, compiled-in
// assets.

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
