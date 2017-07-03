// +build !windows

package ct

import (
	"os"
	"github.com/mattn/go-isatty"
)

func NewCT() ctInterface {
	if os.Getenv("TERM") == "dumb" || !isatty.IsTerminal(os.Stdout.Fd()) {
		return NewDumb()
	}
	return NewAnsi()
}
