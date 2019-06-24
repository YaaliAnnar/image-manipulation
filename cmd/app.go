package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"log"

	"github.com/YaaliAnnar/image-manipulation/imagetext"
	"github.com/golang/freetype"
)

var (
	backgroundPath  = flag.String("background", "../asset/samoyed.jpg", "background image")
	fontPath        = flag.String("fontfile", "../asset/bebas-neue-regular.ttf", "filename of the ttf font")
	outputDirectory = flag.String("outputdir", "./ticket/", "output directory")
	dpi             = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	size            = flag.Float64("size", 144, "font size in points")
)

func main() {
	font, err := imagetext.GetFont(*fontPath)
	if err != nil {
		log.Println(err)
		return
	}

	imageData, err := imagetext.GetImageData(*backgroundPath)
	resultImage := image.NewRGBA(imageData.Bounds())

	context := freetype.NewContext()
	context.SetDPI(*dpi)
	context.SetFontSize(*size)
	context.SetFont(font)
	context.SetSrc(image.Black)

	for i := 0; i < 10; i++ {

		draw.Draw(resultImage, resultImage.Bounds(), imageData, image.Point{0, 0}, draw.Src)
		iString := fmt.Sprintf("%03d", i)

		context.SetDst(resultImage)
		context.SetClip(resultImage.Bounds())

		pt := freetype.Pt(250, 250)
		pt, err = context.DrawString(iString, pt)
		if err != nil {
			log.Println(err)
			continue
		}

		err = imagetext.SaveImage(*outputDirectory+"tiket-"+iString+".png", resultImage)
		if err != nil {
			log.Println(err)
			continue
		}
	}

}
