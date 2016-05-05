package filereaders

import (
	"appengine"
	"encoding/json"
	"github.com/peterhoward42/frost/contract"
	"github.com/peterhoward42/frost/parse"
	"strings"
)

// The type Whitespace is a file reader for whitespace delimited files that can convert the
// content discovered into a JSON representation, using FROSTS conversion rules. It is
// single-shot to avoid the complexity of reinitialising state.
type Whitespace struct {
	jsonData []interface{}
	inputText      string
	requestContext appengine.Context
}

// The function NewWhitespace is the way to instantiate a Whitespace structure and binds the
// instance to particular file contents.
func NewWhitespace(inputText string, ctx appengine.Context) *Whitespace {
	return &Whitespace{
		jsonData:       []interface{}{},
		inputText:      inputText,
		requestContext: ctx,
	}
}

// The Convert() method is the entry point for converting the file contents into JSON.
func (ws *Whitespace) Convert() []byte {
	// Delegate to a sub function for each separate line of input.
	for _, lineTxt := range strings.Split(ws.inputText, "\n") {
		trimmed := strings.TrimSpace(lineTxt)
		if strings.HasPrefix(trimmed, "#") {
			ws.processCommentLine(trimmed)
		} else {
			ws.processNonCommentLine(trimmed)
		}
	}
	// Automated generation of JSON
	// Todo use a constant for the indent string
	theJson, _ := json.MarshalIndent(ws.jsonData, "", "  ")
	return theJson
}

func (ws *Whitespace) processCommentLine(line string) {
}

func (ws *Whitespace) processNonCommentLine(line string) {
	if len(line) == 0 {
		return
	}
	fields := ws.isolateFields(line)
	ws.processLineFields(fields)
}

func (ws *Whitespace) isolateFields(line string) (fields []interface{}) {
	masked := parse.DisguiseDoubleQuotedSegments(line)
	for _, fieldStr := range strings.Fields(masked) {
		fieldStr = parse.UnDisguise(fieldStr)
		fields = append(fields, contract.NewXXXValue(fieldStr))
	}
	return
}

func (ws *Whitespace) processLineFields(fields []interface{}) {
	// Temporarily treat all fields as isolated values.
	ws.jsonData = append(ws.jsonData, fields...)
}
