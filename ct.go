/*
ct package provides functions to change the color of console text.
*/
package ct;

type Color int

const(
    None = Color(iota)
    Black
    Red
    Green
    Yellow
    Blue
    Magenta
    Cyan
    White
)

/*
ResetColor resets the foreground and background to original values
*/
func ResetColor() {
    resetColor()
}

// ChangeColor sets the foreground and background colors. If the value of the color is None, the corresponding color keep unchanged.
// If fgBright or bgBright is set true, corresponding color use bright color.
func ChangeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
    changeColor(fg, fgBright, bg, bgBright)
}
