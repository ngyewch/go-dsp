package plotutils

import (
	"fmt"
	"os"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func SavePlotToFile(p *plot.Plot, path string, width int, height int) error {
	ext := filepath.Ext(path)
	switch ext {
	case ".png":
	default:
		return fmt.Errorf("unsupported file type: %s", ext)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	img := vgimg.New(vg.Length(width)*vg.Inch/vg.Length(vgimg.DefaultDPI), vg.Length(height)*vg.Inch/vg.Length(vgimg.DefaultDPI))
	dc := draw.New(img)
	p.Draw(dc)

	switch ext {
	case ".png":
		png := vgimg.PngCanvas{
			Canvas: img,
		}
		_, err = png.WriteTo(f)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unsupported file type: %s", ext)
	}
}
