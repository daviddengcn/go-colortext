// +build !windows

package ct

import (
	"os"
)

func NewCT() ctInterface {
	if os.Getenv("TERM") == "dumb" {
		return NewDumb()
	}
	return NewAnsi()
}
