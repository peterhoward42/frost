package filereaders

import (
	"appengine"
	"encoding/json"
	"github.com/peterhoward42/frost/contract"
	"github.com/peterhoward42/frost/parse"
	"strings"
)

// The type Whitespace is a file reader for whitespace delimited files that can convert the
// content into a JSON representation, using FROSTS conversion rules. It is bound to a constant
// chunk of input text to avoid the complexity of reinitialising state.
type Whitespace struct {
	/* The jsonData field is the root node container for a sequence of tree-like data
	structure, which is suitable to be converted to json automatically by go's json package. We
	populate it as we parse and convert the input file, using object types like RowOfValues,
	(and others) from frost's contract package.
	*/
	jsonData []interface{}

	// The entire body of input text in a single string including newlines.
	inputText string

	// The input text split into lines and trimmed of leading and trailing whitespace.
	lines []string

	// We capture the context of the http request that launched this calling stack to make it
	// possible to do logging and diagnostics in the app engine world.
	requestContext appengine.Context

	// The parsing process is statefull, insofar as it knows when we are part way through
	// consuming a table. When so, the liveTableSignature holds a non-nil pointer to a
	// RowOfValues object. This defines the type-signature required of all rows in
	// the table.
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

// The Convert() method is the mandate to launch the file contents conversion into JSON. It
// returns the json text produced.
func (ws *Whitespace) Convert() []byte {
	ws.lines = []string{}
	// Read all the lines in first, so the parsing can seek forwards and backwards in the
	// contents to aid analysis.
	for _, line := range strings.Split(ws.inputText, "\n") {
		ws.lines = append(ws.lines, strings.TrimSpace(line))
	}
	// Then simply consume each line in turn, using line-number indexing. (Zero based).
	for i := 0; i < len(ws.lines); i++ {
		ws.processLine(i)
	}
	// Generate and return the JSON.
	theJson, _ := json.MarshalIndent(ws.jsonData, "", "  ")
	return theJson
}

func (ws *Whitespace) processLine(lineIndex int) {
	// This method implements an order of precedence, which is of central significance.
	line := ws.lines[lineIndex]      // Just shorthand
	fields := ws.isolateFields(line) // Includes quoted string magic.

	// We choose not to use the brevity of a switch statement, because we want to use some
	// inline returns, and drop-through behaviour.

	consumed, tableSignature := ws.consumeAsFirstLineInATable(line, fields, lineIndex)
	if consumed {
		ws.liveTableSignature = tableSignature
		return
	}
	if ws.consumeAsTableContinuation(fields) {
		return
	}
	// Single point chosen for when we may assert that we cannot be in the middle of a
	// table.
	ws.liveTableSignature = nil

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

func (ws *Whitespace) isolateFields(line string) (fields []string) {
	lineWithMaskedQuotedStrings := parse.DisguiseDoubleQuotedSegments(line)
	for _, delimitedString := range strings.Fields(lineWithMaskedQuotedStrings) {
		undisguisedString := parse.UnDisguise(delimitedString)
		fields = append(fields, undisguisedString)
	}
	return
}

func (ws *Whitespace) consumeAsFirstLineInATable(
	line string, fields []string, currentLineIndex int) (
	consumed bool, tableSignature *contract.RowOfValues) {
	// Cannot be so if we are already in a table
	if ws.liveTableSignature == nil {
		return false, nil
	}
	// todo, centralise looks like a comment
	// Cannot be so if the line is a comment line
	if strings.HasPrefix(line, "#") {
		return false, nil
	}
	// Cannot be so if the line has fewer than two fields.
	if len(fields) < 2 {
		return false, nil
	}
	// It seems that this line in isolation might be the first row of a table
	tableFirstRow := contract.NewRowOfValues(fields)

	// FROST defines a necessary condition for a table to be that it has at least two rows, and
	// that all rows have matching signatures.
	nextLineIndex := currentLineIndex + 1
	if nextLineIndex >= len(ws.lines) {
		return false, nil
	}
	nextRow := contract.NewRowOfValues(ws.isolateFields(ws.lines[nextLineIndex]))
	if tableFirstRow.HasSameSignatureAs(nextRow) == false {
		return false, nil
	}
	// We conclude the line is the first of a new table.
	return true, tableFirstRow
}

func (ws *Whitespace) consumeAsTableContinuation(fields []string) bool {
	return false
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
