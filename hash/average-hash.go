package hash

import (
	"image"

	"github.com/nfnt/resize"
)

const (
	AImageXSize = 8
	AImageYSize = 8
)

func AHash(img image.Image) uint64 {
	// resize our image
	resImage := resize.Resize(AImageXSize, AImageYSize, img, resize.NearestNeighbor)

	// make grayscale
	grayImage := rgbaToGray(resImage)

	// get our average from the image
	avg := mean(grayImage)

	var val uint8
	var hash uint64
	for y := 0; y < AImageXSize; y++ {
		for x := 0; x < AImageXSize; x++ {
			val = grayImage.GrayAt(x, y).Y
			if val > avg {
				hash |= 1
			}
			hash <<= 1
		}
	}

	// fmt.Printf("%b\n", hash)

	return hash
}

func mean(img *image.Gray) uint8 {
	var sum uint32
	for y := 0; y < AImageXSize; y++ {
		for x := 0; x < AImageXSize; x++ {
			sum += uint32(img.GrayAt(x, y).Y)
		}
	}

	return uint8(sum / (AImageXSize * AImageYSize))
}
