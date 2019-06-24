package imagetext

import (
	"bufio"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Context struct {
	freetype.Context
}

func GetFont(fontPath string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	return font, nil
}

func GetImageData(backgroundPath string) (image.Image, error) {
	imageFile, err := os.Open(backgroundPath)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	imageData, err := jpeg.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}

func SaveImage(outputPath string, image *image.RGBA) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	fileWriter := bufio.NewWriter(outFile)

	err = png.Encode(fileWriter, image)
	if err != nil {
		return err
	}

	err = fileWriter.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (context *Context) AddText(x int, y int, text string) error {
	pt := freetype.Pt(x, y)
	pt, err := context.DrawString(text, pt)
	if err != nil {
		return err
	}
	return nil
}
