package resources

type ExampleInputFileType int // See the enumeration below.

const (
	SpaceDelimitedExample ExampleInputFileType = iota
)

// The GetExampleFileContents function provides convenient access to the canned example
// input files that have been compiled in to the program as assets, in terms of their meta
// role rather than use their file names directly. E.g. "get me the space delimitted input
// file example".
func GetExampleFileContents(exampleType ExampleInputFileType) string {
	switch {
	case exampleType == SpaceDelimitedExample:
		return string(MustAsset(`static/examples/space_delim.txt`))
	default:
		return ""
	}
}
