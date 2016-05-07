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
	jsonData  []interface{}
	inputText string

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
			ws.consumeComment(trimmed)
		} else {
			ws.consumeNonCommentLine(trimmed)
		}
	}
	// Generate the JSON using the json library's ability to traverse automatically the data
	// structure topology we made underneath ws.jsonData.
	theJson, _ := json.MarshalIndent(ws.jsonData, "", "  ")
	return theJson
}

func (ws *Whitespace) consumeComment(line string) {
	ws.jsonData = append(ws.jsonData, contract.NewComment(line))
}

func (ws *Whitespace) consumeNonCommentLine(line string) {
	if len(line) == 0 {
		return
	}
	fieldStrings := ws.isolateFields(line)
	ws.consumeFieldsFromLine(fieldStrings)
}

func (ws *Whitespace) isolateFields(line string) (fields []string) {
	lineWithMaskedQuotedStrings := parse.DisguiseDoubleQuotedSegments(line)
	for _, delimitedString := range strings.Fields(lineWithMaskedQuotedStrings) {
		undisguisedString := parse.UnDisguise(delimitedString)
		fields = append(fields, undisguisedString)
	}
	return
}

func (ws *Whitespace) consumeFieldsFromLine(fields []string) {
	// The order here represents the precedence we declared for recognizing things in
	// the FROST contract for whitespace delimited files.
	switch {
	case ws.consumeAsKeyValuePair(fields):
		return
	case ws.consumeAsStandaloneRowOfValues(fields):
		return
	case ws.consumeAsAsAStandaloneValue(fields):
		return
	default:
		panic("Something has gone wrong. Nothing accepted this input line.")
	}
}

func (ws *Whitespace) consumeAsKeyValuePair(fieldStrings []string) (succeeded bool) {
	// The first field must look like a sensible key
	if parse.LooksLikeAKeyString(fieldStrings[0]) == false {
		return false
	}
	keyString := fieldStrings[0]
	// When there are only two fields, we take the second one to be the value.
	if len(fieldStrings) == 2 {
		ws.jsonData = append(ws.jsonData, contract.NewKeyValuePair(
			keyString, fieldStrings[1]))
		return true
	}
	// We will take the first and third, when there are three, if the middle one is
	// an equals sign or a colon.
	if len(fieldStrings) == 3 {
		if strings.Contains(":=", fieldStrings[1]) {
			ws.jsonData = append(ws.jsonData,
				contract.NewKeyValuePair(keyString, fieldStrings[2]))
			return true
		}
	}
	return false
}

func (ws *Whitespace) consumeAsStandaloneRowOfValues(fieldStrings []string) (succeeded bool) {
	if len(fieldStrings) < 3 {
		return false
	}
	ws.jsonData = append(ws.jsonData, contract.NewRowOfValues(fieldStrings))
	return true
}

func (ws *Whitespace) consumeAsAsAStandaloneValue(fieldStrings []string) (succeeded bool) {
	if len(fieldStrings) > 1 {
		panic("We should only reach here for rows with a single field in.")
	}
	theOnlyFieldString := fieldStrings[0]
	ws.jsonData = append(ws.jsonData, contract.NewXXXValue(theOnlyFieldString))
	return true
}
