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
// bound to a constant chunk of input text to avoid the complexity of reinitialising state.
type Whitespace struct {
	/* As we parse the file contents, we build a data structure topology that will eventually
	be converted to JSON automatically. The jsonData field below is the slice that forms the
	root of this data structure and defines the sequence of objects. We define the types held
	by the slice to be objects that implement the empty interface, because that is what the
	json library will want when it comes to the automatic conversion. The concrete types we
	populate it live in the "contract" package, and have their own native way of expressing
	the data structure topology to be traversed.
	*/
	jsonData []interface{}

	// The entire body of input text in a single string including newlines.
	inputText string

	// The input text split into lines and trimmed of leading and trailing whitespace.
	lines []string

	// We capture the context of the http request that launched this calling stack to make it
	// possible to do logging and diagnostics in the app engine world.
	requestContext appengine.Context

	liveTableSignature *contract.RowOfValues
}

// The function NewWhitespace() is the way to instantiate a Whitespace structure and binds this
// instance permanently to a particular file contents, and a single http request instance.
func NewWhitespace(inputText string, ctx appengine.Context) *Whitespace {
	return &Whitespace{
		inputText:          inputText,
		lines:              nil,
		requestContext:     ctx,
		liveTableSignature: nil,
	}
}

// The Convert() method is the mandate to launch the file contents conversion into JSON.
func (ws *Whitespace) Convert() []byte {
	ws.lines = []string{}
	for _, line := range strings.Split(ws.inputText, "\n") {
		ws.lines = append(ws.lines, strings.TrimSpace(line))
	}
	for i := 0; i < len(ws.lines); i++ {
		ws.processLine(i)
	}
	// Generate the JSON using the json library's ability to traverse automatically the data
	// structure topology we made underneath ws.jsonData.
	theJson, _ := json.MarshalIndent(ws.jsonData, "", "  ")
	return theJson
}

func (ws *Whitespace) processLine(lineIndex int) {
	// This method implements a highly significant order of precedence.
	line := ws.lines[lineIndex]
	fields := ws.isolateFields(line)

	// We choose not to use the brevity of a switch statement, because we want to use some
	// inline returns, and drop-through behaviour.

	tableSignature, consumed := ws.consumeAsFirstLineInATable(fields, lineIndex)
	if consumed {
		ws.liveTableSignature = tableSignature
		return
	}
	if ws.consumeAsTableContinuation(fields) {
		return
	}
	// Single point at which to observe implicitly that we have fallen off the end of a
	// table under construction.
	if ws.liveTableSignature != nil {
		ws.liveTableSignature = nil
	}
	if ws.consumeAsEmptyLine(line) {
		return

	}
	if ws.consumeAsCommentLine(line) {
		return
	}
	if ws.consumeAsKeyValuePair(fields) {
		return
	}
	if ws.consumeAsStandaloneRowOfValues(fields) {
		return
	}
	if ws.consumeAsAsAStandaloneValue(fields) {
		return
	}
	panic("Something has gone wrong. Nothing accepted this input line.")
}

func (ws *Whitespace) consumeAsCommentLine(line string) bool {
	if strings.HasPrefix(line, "#") {
		ws.jsonData = append(ws.jsonData, contract.NewComment(line))
		return true
	}
	return false
}

func (ws *Whitespace) consumeAsEmptyLine(line string) bool {
	return len(line) == 0
}

func (ws *Whitespace) isolateFields(line string) (fields []string) {
	lineWithMaskedQuotedStrings := parse.DisguiseDoubleQuotedSegments(line)
	for _, delimitedString := range strings.Fields(lineWithMaskedQuotedStrings) {
		undisguisedString := parse.UnDisguise(delimitedString)
		fields = append(fields, undisguisedString)
	}
	return
}

func (ws *Whitespace) consumeAsFirstLineInATable(
	fields []string, currentLineIndex int) (
	tableSignature *contract.RowOfValues, consumed bool) {
	return nil, false
}

func (ws *Whitespace) consumeAsTableContinuation(fields []string) bool {
	return false
}

func (ws *Whitespace) consumeAsKeyValuePair(fields []string) (succeeded bool) {
	// The first field must look like a sensible key
	if parse.LooksLikeAKeyString(fields[0]) == false {
		return false
	}
	keyString := fields[0]
	// When there are only two fields, we take the second one to be the value.
	if len(fields) == 2 {
		ws.jsonData = append(ws.jsonData, contract.NewKeyValuePair(
			keyString, fields[1]))
		return true
	}
	// We will take the first and third, when there are three, if the middle one is
	// an equals sign or a colon.
	if len(fields) == 3 {
		if strings.Contains(":=", fields[1]) {
			ws.jsonData = append(ws.jsonData,
				contract.NewKeyValuePair(keyString, fields[2]))
			return true
		}
	}
	return false
}

func (ws *Whitespace) consumeAsStandaloneRowOfValues(fields []string) (succeeded bool) {
	if len(fields) < 3 {
		return false
	}
	ws.jsonData = append(ws.jsonData, contract.NewRowOfValues(fields))
	return true
}

func (ws *Whitespace) consumeAsAsAStandaloneValue(fields []string) (succeeded bool) {
	if len(fields) > 1 {
		panic("We should only reach here for rows with a single field in.")
	}
	theOnlyFieldString := fields[0]
	ws.jsonData = append(ws.jsonData, contract.NewXXXValue(theOnlyFieldString))
	return true
}
