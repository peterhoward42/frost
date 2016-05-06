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
	// As we parse the file contents, we build a data structure topology that will eventually
	// be converted to JSON automatically. The jsonData field below is the slice that forms the
	// root of this data structure. We define the contained type to be objects that implement
	// the empty interface, because that is what the json library will want when it comes to
	// the automatic conversion. The concrete types we populate it live in the "contract"
	// package, and have their own native way of expressing the data structure topology to be
	// traversed.
	jsonData []interface{}
	inputText      string

	// We capture the context of the http request that launched this calling stack to make it
	// possible to do logging and diagnostics in the app engine world.
	requestContext appengine.Context
}

// The function NewWhitespace() is the way to instantiate a Whitespace structure and binds this
// instance permanently to a particular file contents, and a single http request instance.
func NewWhitespace(inputText string, ctx appengine.Context) *Whitespace {
	return &Whitespace{
		jsonData:       []interface{}{},
		inputText:      inputText,
		requestContext: ctx,
	}
}

// The Convert() method is the mandate to launch the file contents conversion into JSON.
func (ws *Whitespace) Convert() []byte {
	// In the case of whitespace delimited files, we delegate to a sub function for each
	// separate line of input.
	for _, lineTxt := range strings.Split(ws.inputText, "\n") {
		trimmed := strings.TrimSpace(lineTxt)
		if strings.HasPrefix(trimmed, "#") {
			ws.processCommentLine(trimmed)
		} else {
			ws.processNonCommentLine(trimmed)
		}
	}
	// Generate the JSON using the json library's ability to traverse automatically the data
	// structure topology we made underneath ws.jsonData.
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
