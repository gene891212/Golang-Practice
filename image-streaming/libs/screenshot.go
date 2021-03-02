package libs

import (
	"bytes"
	"image/png"

	"github.com/kbinani/screenshot"
)

// GetScreenshot ...
// func GetScreenshot() (*bytes.Buffer, error) {
// 	bounds := screenshot.GetDisplayBounds(1)
// 	buff := new(bytes.Buffer)
// 	img, err := screenshot.CaptureRect(bounds)
// 	// fileName := fmt.Sprintf("img/%s.png", time.Now().String())
// 	// file, _ := os.Create(fileName)
// 	// defer file.Close()
// 	// png.Encode(file, img)
// 	// png.Encode(buff, img)
// 	jpeg.Encode(buff, img, nil)
// 	return buff, err
// }

func GetScreenshot() (*bytes.Buffer, error) {
	bounds := screenshot.GetDisplayBounds(1)
	buff := new(bytes.Buffer)
	img, err := screenshot.CaptureRect(bounds)
	png.Encode(buff, img)
	return buff, err
}
