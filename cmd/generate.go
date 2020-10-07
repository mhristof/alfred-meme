package cmd

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/spf13/cobra"
)

const fontSize = 36
const strokeSize = 6

var (
	generateCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Take over the world",
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]

			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error, file %s does not exist [%s]", file, err)
				os.Exit(1)
			}
			defer f.Close()

			text, err := cmd.Flags().GetString("text")
			if err != nil {
				panic(err)
			}

			lines := strings.Split(strings.TrimSuffix(strings.TrimPrefix(text, `"`), `"`), "|")

			meme, _, err := image.Decode(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error, cannot load image %s [%s]", file, err)
				os.Exit(1)
			}

			updateImage(meme, lines)

		},
	}
)

func updateImage(img image.Image, text []string) {
	r := img.Bounds()
	w := r.Dx()
	h := r.Dy()

	m := gg.NewContext(w, h)
	m.DrawImage(img, 0, 0)
	m.LoadFontFace("/System/Library/Fonts/Supplemental/Comic Sans MS.ttf", fontSize)

	// Apply black stroke
	m.SetHexColor("#000")
	//lastx, lasty := 0.0, 0.0
	m = drawText(m, text[0], float64(w/2), float64(strokeSize), float64(w)/2, float64(fontSize+strokeSize))
	if len(text) == 2 {
		m = drawText(m, text[1], float64(w/2), float64(h-fontSize*2), float64(w)/2, float64(h)-fontSize)
	}

	m.SavePNG("meme.jpg")
	fmt.Println("meme.jpg")
}

func drawText(m *gg.Context, text string, startX, startY, anchoredX, anchoredY float64) *gg.Context {
	m.SetHexColor("#000")
	for dy := -strokeSize; dy <= strokeSize; dy++ {
		for dx := -strokeSize; dx <= strokeSize; dx++ {
			// give it rounded corners
			if dx*dx+dy*dy >= strokeSize*strokeSize {
				continue
			}
			x := startX + float64(dx)
			y := startY + float64(fontSize+dy)
			m.DrawStringAnchored(text, x, y, 0.5, 0.5)
		}
	}

	// Apply white fill
	m.SetHexColor("#FFF")
	m.DrawStringAnchored(text, anchoredX, anchoredY, 0.5, 0.5)
	return m
}

func init() {
	generateCmd.PersistentFlags().StringP("text", "t", "", "Text for the meme with '|' as the new line separator")
	rootCmd.AddCommand(generateCmd)
}
