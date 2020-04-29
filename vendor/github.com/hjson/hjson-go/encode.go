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

var needsEscape, needsQuotes, needsEscapeML