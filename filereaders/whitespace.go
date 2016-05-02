package filereaders

import (
	"strings"
	"parse"
	"contract"
)

// The type Whitespace is a file reader for whitespace delimited files that can convert the
// content discovered into a JSON representation, using FROSTS conversion rules. It is
// shingle-shot to avoid the complexity of reinitialising state.
type Whitespace struct {
	Json []byte
	inputText string
	parsingContext *parse.Context	// Holds the original input line for error reporting needs.
}

// The function NewWhitespace is the way to instantiate a Whitespace structure and binds the
// instance to particular file contents.
func NewWhitespace(inputText string) *Whitespace {
	return &Whitespace {
		Json: []byte{},
		inputText: inputText,
		parsingContext: nil,
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
	return ws.Json
}

func (ws *Whitespace) processCommentLine(line string) {
}

func (ws *Whitespace) processNonCommentLine(line string) {
	if len(line) == 0 {
		return
	}
	fields := ws.isolateFields(line);
	ws.processFields(fields);
}

func (ws *Whitespace) isolateFields(line string) []string {
	line = parse.MaskDoubleQuotes(line)
	fields := strings.Fields(line)
	for idx, field := range(fields) {
		fields[idx] = parse.UnMaskDoubleQuotes(field);
	}
	return fields
}

func (ws *Whitespace) processFields(fields []string) {
	for _, field := range(fields) {
		valueAlone := contract.NewValueAlone(field)
		ws.Json = append(ws.Json, valueAlone.Json...);
	}
}