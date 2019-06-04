package hash

import (
	"image"

	"github.com/nfnt/resize"
)

const (
	DImageXSize = 9
	DImageYSize = 8
)

func DHash(img image.Image) uint64 {
	// resize our image
	resImage := resize.Resize(DImageXSize, DImageYSize, img, resize.NearestNeighbor)

	// make grayscale
	grayImage := rgbaToGray(resImage)

	// loop over our image, and set hash accordingly
	var hash uint64
	var left, right uint8

	for y := 0; y < DImageYSize; y++ {
		// set our initial left value
		left = grayImage.GrayAt(0, y).Y
		for x := 1; x < DImageXSize; x++ {
			right = grayImage.GrayAt(x, y).Y
			if right < left {
				hash |= 1
			}
			hash <<= 1
		}
	}

	// fmt.Printf("%064b\n", hash)

	return hash
}
