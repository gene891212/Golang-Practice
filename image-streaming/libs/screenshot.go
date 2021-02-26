package libs

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

// GetScreenshot ...
func GetScreenshot() *image.RGBA {
	bounds := screenshot.GetDisplayBounds(1)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	fileName := "screen.png"
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)

	fmt.Println("Finish")
	return img
}
