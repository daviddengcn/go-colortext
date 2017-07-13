package ct

import (
	"fmt"
	"strconv"
)

func ansiText(fg Color, fgBright bool, bg Color, bgBright bool) string {
	if fg == None && bg == None {
		return ""
	}
	s := []byte("\x1b[0")
	if fg != None {
		s = strconv.AppendUint(append(s, ";"...), 30+(uint64)(fg-Black), 10)
		if fgBright {
			s = append(s, ";1"...)
		}
	}
	if bg != None {
		s = strconv.AppendUint(append(s, ";"...), 40+(uint64)(bg-Black), 10)
		if bgBright {
			s = append(s, ";1"...)
		}
	}
	s = append(s, "m"...)
	return string(s)
}

type ctAnsi struct {
} 

func NewAnsi() *ctAnsi {
	return &ctAnsi{
	}
}
func (c *ctAnsi) resetColor() {
	fmt.Print("\x1b[0m")
}

func (c *ctAnsi) changeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
	if fg == None && bg == None {
		return
	}
	fmt.Print(ansiText(fg, fgBright, bg, bgBright))
}
