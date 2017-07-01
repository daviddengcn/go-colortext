// +build !windows

package ct

import (
	"os"
	"github.com/mattn/go-isatty"
)

func GetTerminal() TermType {
	if os.Getenv("TERM") == "dumb" || !isatty.IsTerminal(os.Stdout.Fd()){
		return DumbTerm
	}
	return AnsiTerm
}

type ctWin struct {
} 

func NewWin() *ctWin {
	return &ctWin{
	}
}

func (c *ctWin) resetColor() {
}

func (c *ctWin) changeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
}
