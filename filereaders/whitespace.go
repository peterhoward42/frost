package filereaders

import (
	"encoding/json"
	"github.com/peterhoward42/frost/contract"
	"github.com/peterhoward42/frost/parse"
	"strings"
)

// The type WhitespaceConverter is a file reader for whitespace delimited files that can convert the
// content into a JSON representation, using FROSTS conversion rules. It is bound to a constant
// chunk of input text to avoid the complexity of reinitialising state.
type WhitespaceConverter struct {
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

	// When our parsing reaches a table, this field is both a handle to the table being
	// augmented, and a signal when non null of this state existing.
	tableUnderConstruction *contract.Table
}

// The function NewWhitespaceConverter() is the way to instantiate a Whitespace structure and binds this
// instance permanently to a particular file contents, and a single http request instance.
func NewWhitespaceConverter(inputText string) *WhitespaceConverter {
	return &WhitespaceConverter{
		inputText: inputText,
		lines:     nil,
		tableUnderConstruction: nil,
	}
}

// The Convert() method is the mandate to launch the file contents conversion into JSON. It
// returns the json text produced.
func (ws *WhitespaceConverter) Convert() []byte {
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

func (ws *WhitespaceConverter) processLine(lineIndex int) {
	line := ws.lines[lineIndex]      // Just shorthand
	fields := ws.isolateFields(line) // Includes quoted string magic.

	// This expresses our wish to accept the first of these that succeeds, and then to
	// stop.

	// The first two cases are about detecting the start, the continuation and falling off
	// the end of a table. They come first in the order that they do, as part of managing and
	// holding this state.
	switch {
	case ws.consumeAsFirstLineInATable(line, fields, lineIndex):
	case ws.consumeAsTableContinuation(line, fields):
	case ws.consumeAsEmptyLine(line):
	case ws.consumeAsCommentLine(line):
	case ws.consumeAsKeyValuePair(fields):
	case ws.consumeAsStandaloneRowOfValues(fields):
	case ws.consumeAsAsAStandaloneValue(fields):
	default:
		panic("Something has gone wrong. Nothing accepted this input line.")
	}
}

func (ws *WhitespaceConverter) isAComment(line string) bool {
	return strings.HasPrefix(line, "#")
}

func (ws *WhitespaceConverter) isolateFields(line string) (fields []string) {
	lineWithMaskedQuotedStrings := parse.DisguiseDoubleQuotedSegments(line)
	for _, delimitedString := range strings.Fields(lineWithMaskedQuotedStrings) {
		undisguisedString := parse.UnDisguise(delimitedString)
		fields = append(fields, undisguisedString)
	}
	return
}

func (ws *WhitespaceConverter) consumeAsFirstLineInATable(
	line string, fields []string, currentLineIndex int) (consumed bool) {
	// Cannot be so if we are already in a table
	if ws.tableUnderConstruction != nil {
		return false
	}
	if ws.isAComment(line) {
		return false
	}
	// Cannot be so if the line has fewer than two fields.
	if len(fields) < 2 {
		return false
	}
	// It seems that this line in isolation might be the first row of a table
	tableFirstRow := contract.NewRowOfValues(fields)

	// FROST defines a necessary condition for a table to be that it has at least two rows, and
	// that all rows have matching signatures.
	nextLineIndex := currentLineIndex + 1
	if nextLineIndex >= len(ws.lines) {
		return false
	}
	nextRow := contract.NewRowOfValues(ws.isolateFields(ws.lines[nextLineIndex]))
	if tableFirstRow.HasSameSignatureAs(nextRow) == false {
		return false
	}
	// We conclude the line is the first of a new table.
	ws.tableUnderConstruction = contract.NewTable(tableFirstRow)
	ws.jsonData = append(ws.jsonData, ws.tableUnderConstruction)
	return true
}

func (ws *WhitespaceConverter) consumeAsTableContinuation(line string, fields []string) bool {
	if ws.tableUnderConstruction == nil {
		return false
	}
	if ws.isAComment(line) {
		ws.tableUnderConstruction = nil
		return false
	}
	tableRow := contract.NewRowOfValues(fields)
	if tableRow.HasSameSignatureAs(ws.tableUnderConstruction.Signature) {
		ws.tableUnderConstruction.AddRow(tableRow)
		return true
	}
	ws.tableUnderConstruction = nil
	return false
}

func (ws *WhitespaceConverter) consumeAsCommentLine(line string) bool {
	if ws.isAComment(line) {
		ws.jsonData = append(ws.jsonData, contract.NewComment(line))
		return true
	}
	return false
}

func (ws *WhitespaceConverter) consumeAsEmptyLine(line string) bool {
	return len(line) == 0
}

func (ws *WhitespaceConverter) consumeAsKeyValuePair(fields []string) (succeeded bool) {
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

func (ws *WhitespaceConverter) consumeAsStandaloneRowOfValues(fields []string) (succeeded bool) {
	if len(fields) < 3 {
		return false
	}
	ws.jsonData = append(ws.jsonData, contract.NewRowOfValues(fields))
	return true
}

func (ws *WhitespaceConverter) consumeAsAsAStandaloneValue(fields []string) (succeeded bool) {
	if len(fields) > 1 {
		panic("We should only reach here for rows with a single field in.")
	}
	theOnlyFieldString := fields[0]
	ws.jsonData = append(ws.jsonData, contract.NewXXXValue(theOnlyFieldString))
	return true
}
