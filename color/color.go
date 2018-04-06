package color

import (
	"fmt"
)

const (
	clearScreen = "\033[2J\033[H"
)

// Printer handles printing colorized text
type Printer interface {
	Print(s string) string
}

type printer struct {
	code int
}

const (
	startColorCode = 31
	endColorCode   = 37
)

// New creates a new Printer that colorized text,
// cycling through a fixed color palette
func New() Printer {
	return &printer{
		code: startColorCode,
	}
}

func (p *printer) Print(s string) string {
	val := fmt.Sprintf("%s\033[0;%dm%s\033[0m\n", clearScreen, p.code, s)
	p.code++
	if p.code > endColorCode {
		p.code = startColorCode
	}
	return val
}
