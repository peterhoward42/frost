package parse

type Context struct {
	LineNumber   int
	OriginalLine string
}

func NewContext(lineNumber int, originalLine string) *Context {
	return &Context{
		LineNumber:   lineNumber,
		OriginalLine: originalLine,
	}
}
