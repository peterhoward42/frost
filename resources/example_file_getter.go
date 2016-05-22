package resources

type ExampleInputFileType int // See the enumeration below.

const (
	SpaceDelimitedExample ExampleInputFileType = iota
)

func GetExampleFileContents(exampleType ExampleInputFileType) string {
	switch {
	case exampleType == SpaceDelimitedExample:
		return string(MustAsset(`static/examples/space_delim.txt`))
	default:
		return ""
	}
}
