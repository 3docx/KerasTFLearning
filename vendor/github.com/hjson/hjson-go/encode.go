package hjson

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// EncoderOptions defines options for encoding to Hjson.
type EncoderOptions struct {
	// End of line, should be either \n or \r\n
	Eol string
	// Place braces on the same line
	BracesSameLine bool
	// Deprecated: Hjson always emits braces
	EmitRootBraces bool
	// Always place string in quotes
	QuoteAlways bool
	// Indent string
	IndentBy string
	// Allow the -0 value (unlike ES6)
	AllowMinusZero bool
	// Encode unknown values as 'null'
	UnknownAsNull bool
}

// DefaultOptions returns the default encoding options.
func DefaultOptions() EncoderOptions {
	opt := EncoderOptions{}
	opt.Eol = "\n"
	opt.BracesSameLine = false
	opt.EmitRootBraces = true
	opt.QuoteAlways = false
	opt.IndentBy = "  "
	opt.AllowMinusZero = false
	opt.UnknownAsNull = false
	return opt
}

type hjsonEncoder struct {
	bytes.Buffer // output
	EncoderOptions
	indent int
}

var needsEscape, needsQuotes, needsEscapeML, startsWithKeyword, needsEscapeName *regexp.Regexp

func init() {
	var commonRange = `\x7f-\x9f\x{00ad}\x{0600}-\x{0604}\x{070f}\x{17b4}\x{17b5}\x{200c}-\x{200f}\x{2028}-\x{202f}\x{2060}-\x{206f}\x{feff}\x{fff0}-\x{ffff}`
	// needsEscape tests if the string can be written without escapes
	needsEscape = regexp.MustCompile(`[\\\"\x00-\x1f` + commonRange + `]`)
	// needsQuotes tests if the string can be written as a quoteless string (includes needsEscape but without \\ and \")
	needsQuotes = regexp.MustCompile(`^\s|^"|^'|^#|^/\*|^//|^\{|^\}|^\[|^\]|^:|^,|\s$|[\x00-\x1f\x7f-\x9f\x{00ad}\x{0600}-\x{0604}\x{070f}\x{17b4}\x{17b5}\x{200c}-\x{200f}\x{2028}-\x{202f}\x{2060}-\x{206f}\x{feff}\x{fff0}-\x{ffff}]`)
	// needsEscapeML tests if the string can be written as a multiline string (like needsEscape but without \n, \r, \\, \", \t)
	var x08Or9 = `\x08` // `\x09` for the old behavior
	needsEscapeML = regexp.MustCompile(`'''|^[\s]+$|[\x00-` + x08Or9 + `\x0b\x0c\x0e-\x1f` + commonRange + `]`)
	// starts with a keyword and optionally is followed by a comment
	startsWithKeyword = regexp.MustCompile(`^(true|false|null)\s*((,|\]|\}|#|//|/\*).*)?$`)
	needsEscapeName = regexp.MustCompile(`[,\{\[\}\]\s:#"']|//|/\*`)
}

var meta = map[byte][]byte{
	// table of character substitutions
	'\b': []byte("\\b"),
	'\t': []byte("\\t"),
	'\n': []byte("\\n"),
	'\f': []byte("\\f"),
	'\r': []byte("\\r"),
	'"':  []byte("\\\""),
	'\\': []byte("\\\\"),
}

func (e *hjsonEncoder) quoteReplace(text string) string {
	return string(needsEscape.ReplaceAllFunc([]byte(text), func(a []byte) []byte {
		c := meta[a[0]]
		if c != nil {
			return c
		}
		return []byte(fmt.Sprintf("\\u%04x", c))
	}))
}

func (e *hjsonEncoder) quote(value string, separator string, isRootObject bool) {

	// Check if we can insert this string without quotes
	// see hjson syntax (must not parse as true, false, null or number)

	if len(value) == 0 {
		e.WriteString(separator + `""`)
	} else if e.QuoteAlways ||
		needsQuotes.MatchString(value) ||
		startsWithNumber([]byte(value)) ||
		startsWithKeyword.MatchString(value) {

		// If the string contains no control characters, no quote characters, and no
		// backslash characters, then we can safely slap some quotes around it.
		// Otherwise we first check if the string can be expressed in multiline
		// format or we must replace the offending characters with safe escape
		// sequences.

		if !needsEscape.MatchString(value) {
			e.WriteString(separator + `"` + value + `"`)
		} else if !needsEscapeML.MatchString(value) && !isRootObject {
			e.mlString(value, separator)
		} else {
			e.WriteString(separator + `"` + e.quoteReplace(value) + `"`)
		}
	} else {
		// return without quotes
		e.WriteString(separator + value)
	}
}

func (e *hjsonEncoder) mlString(value string, separator string) {
	// wrap the string into the ''' (multiline) format

	a := strings.Split(strings.Replace(value, "\r", "", -1), "\n")

	if len(a) == 1 {
		// The string contains only a single line. We still use the multiline
		// format as it avoids escaping the \ character (e.g. when used in a
		// regex).
		e.WriteString(separator + "'''")
		e.WriteString(a[0])
	} else {
		e.writeIndent(e.indent + 1)
		e.WriteString("'''")
		for _, v := range a {
			indent := e.indent + 1
			if len(v) == 0 {
				indent = 0
			}
			e.writeIndent(indent)
			e.WriteString(v)
		}
		e.writeIndent(e.indent + 1)
	}
	e.WriteString("'''")
}

func (e *hjsonEncoder) quoteName(name string) string {
	if len(name) == 0 {
		return `""`
	}

	// Check if we can insert this name without quotes

	if needsEscapeName.MatchString(name) {
		if needsEscape.MatchString(name) {
			name = e.quoteReplace(name)
		}
		return `"` + name + `"`
	}
	// without quotes
	return name
}

type sortAlpha []reflect.Value

func (s sortAlpha) Len() int {
	return len(s)
}
func (s sortAlpha) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortAlpha) Less(i, j int) bool {
	return s[i].String() 