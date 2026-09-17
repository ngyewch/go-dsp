package plot

import (
	"image/color"

	"github.com/dim13/colormap"
	"gonum.org/v1/plot/palette"
)

var (
	defaultSpectrogramPalette = ToPalette(colormap.Inferno)
)

type simplePalette []color.Color

func (palette simplePalette) Colors() []color.Color {
	return palette
}

func ToPalette(colors []color.Color) palette.Palette {
	return simplePalette(colors)
}
