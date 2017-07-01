package ct

type ctDumb struct {
} 

func NewDumb() *ctDumb {
	return &ctDumb{
	}
}

func (c *ctDumb) resetColor() {
}

func (c *ctDumb) changeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
}
