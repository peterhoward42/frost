package filereaders

import (
	"appengine"
	"contract"
	"encoding/json"
	"parse"
	"strings"
)

// The type Whitespace is a file reader for whitespace delimited files that can convert the
// content discovered into a JSON representation, using FROSTS conversion rules. It is
// shingle-shot to avoid the complexity of reinitialising state.
type Whitespace struct {
	JsonData       []contract.JsonType
	inputText      string
	requestContext appengine.Context
	parsingContext *parse.Context // Holds the original input line for error reporting needs.
}

// The function NewWhitespace is the way to instantiate a Whitespace structure and binds the
// instance to particular file contents.
func NewWhitespace(inputText string, ctx appengine.Context) *Whitespace {
	return &Whitespace{
		JsonData:       []contract.JsonType{},
		inputText:      inputText,
		requestContext: ctx,
		parsingContext: nil,
	}
}

// The Convert() method is the entry point for converting the file contents provided to the
// constructor into the FROST-converted JSON structure.
func (ws *Whitespace) Convert() []byte {
	// Delegate to a sub function for each separate line of input.
	for idx, lineTxt := range strings.Split(ws.inputText, "\n") {
		ws.parsingContext = parse.NewContext(idx+1, lineTxt)
		trimmed := strings.TrimSpace(lineTxt)
		if strings.HasPrefix(trimmed, "#") {
			ws.processCommentLine(trimmed)
		} else {
			ws.processNonCommentLine(trimmed)
		}
	}
	// Automated generation of JSON
	// Todo use a constant for the indent string
	theJson, _ := json.MarshalIndent(ws.JsonData, "", "  ")
	return theJson
}

func (ws *Whitespace) processCommentLine(line string) {
}

func (ws *Whitespace) processNonCommentLine(line string) {
	if len(line) == 0 {
		return
	}
	ws.requestContext.Infof("XXX Line is: %v", line)
	fields := ws.isolateFields(line)
	ws.processFields(fields)
}

func (ws *Whitespace) isolateFields(line string) (fields []contract.JsonType) {
	masked := parse.MaskDoubleQuotes(line)
	for _, fieldStr := range strings.Fields(masked) {
		fieldStr = parse.UnMaskDoubleQuotes(fieldStr)
		ws.requestContext.Infof("XXX dFrag is: %v", fieldStr)
		fields = append(fields, contract.NewValueAlone(fieldStr))
	}
	return
}

func (ws *Whitespace) processFields(fields []contract.JsonType) {
	// Temporarily treat all fields as isolated values.
	ws.JsonData = append(ws.JsonData, fields...)
}
