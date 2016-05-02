package filereaders

import (
	"strings"
	"parse"
)

// The type Whitespace is a file reader for whitespace delimited files that can convert the
// content discovered into a JSON representation, using FROSTS conversion rules. It is
// shingle-shot to avoid the complexity of reinitialising state.
type Whitespace struct {
	inputText string
	context *parse.Context	// Holds the original input line for error reporting needs.
}

// The function NewWhitespace is the way to instantiate a Whitespace structure and binds the
// instance to particular file contents.
func NewWhitespace(inputText string) *Whitespace {
	return &Whitespace {
		inputText: inputText,
	}
}

// The Convert() method is the entry point for converting the file contents provided to the
// constructor into the FROST-converted JSON structure.
func (ws *Whitespace) Convert() []byte {
	// Delegate to a sub function for each separate line of input.
	for idx, lineTxt := range(strings.Split(ws.inputText, "\n")) {
		lineNumber := idx + 1
		ws.context = parse.NewContext(lineNumber, lineTxt)
		trimmed := strings.TrimSpace(lineTxt);
		if strings.HasPrefix(trimmed, "#") {
			ws.processCommentLine(trimmed)
		} else {
			ws.processNonCommentLine(trimmed)
		}
	}
	return []byte("will be whitespace content")
}

func (ws *Whitespace) processCommentLine(line string) {
}

func (ws *Whitespace) processNonCommentLine(line string) {
}