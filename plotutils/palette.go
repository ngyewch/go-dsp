package plotutils

import (
	"image/color"

	"gonum.org/v1/plot/palette"
)

type simplePalette []color.Color

func (palette simplePalette) Colors() []color.Color {
	return palette
}

func PaletteFromColors(colors []color.Color) palette.Palette {
	return simplePalette(colors)
}
