package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

// SpandIMG is an image in spand
type SpandIMG struct {
	Width, Height int
	Path          string
	src           image.Image
}

func (img image.Image) setPixel(x, y, c int) {

}

func main() {

	// Read image from file that already exists
	imgFile, err := os.Open("moto.png")
	if err != nil {
		// Handle error
	}
	defer imgFile.Close()

	// Metadata
	config, format, err := image.DecodeConfig(bufio.NewReader(imgFile))
	if err != nil {
		log.Fatal(err) // handle image unknown format and use convertopng
	}
	fmt.Println(format, config.Height, config.Width, config.ColorModel)

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	imgFile.Seek(0, 0)
	imgData, _, err := image.Decode(imgFile)
	//size := imgData.Bounds().Size()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(imgData)
	//fmt.Println(imgType)

	//img := SpandIMG{size.X, size.Y, "moto.png", imgData}

	outFile, err := os.Create("changed.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	png.Encode(outFile, imgData)

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	//loadedImage, err := png.Decode(existingImageFile)
	//if err != nil {
	// Handle error
	//}
	//fmt.Println(loadedImage)
}
