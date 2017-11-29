// +build !windows

package ct

import (
	"os"
	"github.com/mattn/go-isatty"
)

func resetColor() {
	ansiResetColor()
}

func changeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
	ansiChangeColor(fg, fgBright, bg, bgBright)
}

func writerIsTerminal() bool {
	if v, ok := Writer.(*os.File); ok {
		if isatty.IsTerminal(v.Fd()) {
			return true
		}
	}
	return false
}