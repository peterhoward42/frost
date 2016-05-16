package resources

import (
	"github.com/elazarl/go-bindata-assetfs"
	"net/http"
)

var CompiledFileSystem http.FileSystem =
		&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: ""}
