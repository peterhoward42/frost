package resources

import (
	"github.com/elazarl/go-bindata-assetfs"
	"net/http"
)

// The compiledFileSystem package-global variable is an object which looks to the outside world like
// an http.FileSystem, but which under the hood is backed by resources captured at compile time.
// It is exported for the benefit of initialising the web server at boot time, but is not intended
// to be accessed from outside the package for any other reason. Instead a set of conceptually,
// higher level apis are provided. For example GetExampleInputFile().
var CompiledFileSystem http.FileSystem = &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir}
