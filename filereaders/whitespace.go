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
	parsingContext *parse.Context	// Holds the original input line for error reporting needs.
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
		ws.parsingContext = parse.NewContext(idx + 1, lineTxt)
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
	if len(line) == 0 {
		return
	}
	fields := isolateFields(line);
	processFields(fields);
}

func (ws *Whitespace) isolateFields(line string) []string {
	line = parse.MaskDoubleQuotes(line)
	fields := strings.Fields(line)
	for idx, field := range(fields) {
		fields[idx] = parse.UnMaskDoubleQuotes(field);
	}
	return fields
}