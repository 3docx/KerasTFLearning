package hjson

import (
	"errors"
	"math"
	"strconv"
)

type parseNumber struct {
	data []byte
	at   int  // The index of the current character
	ch   byte // The current character
}

func (p *parseNumber) next() bool {
	// get the next character.
	len := len(p.data)
	if p.at < len {
		p.ch = p.data[p.at]
		p.at++
		return true
	}
	if p.at == len {
		p.at++
		p.ch = 0
	}
	return false
}

func startsWithNumber(text []byte) bool {
	if _, err := tryParseNumber(text, true); err == nil {
		return true
	}
	return false
}

func tryParseNumber(text []byte, stopAtNext bool) (float64, error) {
	// Parse a number value.

	p := parseNumber{text, 0, ' '}
	leadingZeros := 0
	testLeading := true
	p.next()
	if p.ch == '-' {
		p.next()
	}
	for p.ch >= '0' && p.ch <= '9' {
		if testLeading {
			if p.ch == '0' {
				leadingZeros++
			} else {
				testLeading = false
			}
		}
		p.next()
	}
	if testLeading {
		leadingZeros--
	} // single 0 is allowed
	if p.ch == '.' {
		for p.next() && p.ch >= '0' && p.ch <= '9' {
		}
	}
	if p.ch == 'e' || p.ch == 'E' {
		p.next()
		if p.ch == '-' || p.ch == '+' {
			p.next()
		}
		for p.ch >= '0' && p.ch <= '9' {
			p.next()
		}
	}

	end := p.at

	// skip white/to (newline)
	for p.ch > 0 && p.ch <= ' ' {
		p.next()
	}

	if stopAtNext {
		// end scan if we find a punctuator character like ,}] or a comment
		if p.ch == ',' || p.ch == '}' || p.ch == ']' ||
			p.ch == '#' || p.ch == '/' && (p.data[p.at] == '/' || p.data[