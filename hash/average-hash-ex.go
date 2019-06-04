package hash

import (
	"image"

	"github.com/nfnt/resize"
)

const (
	AEImageXSize = 16
	AEImageYSize = 16
)

func AEHash(img image.Image) [4]uint64 {
	// resize our image
	resImage := resize.Resize(AEImageXSize, AEImageYSize, img, resize.NearestNeighbor)

	// make grayscale
	grayImage := rgbaToGray(resImage)

	// get our average from the image
	avg := mean2(grayImage)

	var val uint8
	hash := [4]uint64{}
	var currentIndex uint8
	for y := 0; y < AEImageXSize; y++ {
		for x := 0; x < AEImageXSize; x++ {
			val = grayImage.GrayAt(x, y).Y
			if val > avg {
				hash[currentIndex] |= 1
			}
			hash[currentIndex] <<= 1

			if ((y*AEImageYSize)+1+x)%64 == 0 {
				currentIndex++
				// fmt.Println("changed..")
			}
		}
	}

	// fmt.Printf("%b\n", hash)

	return hash
}

func mean2(img *image.Gray) uint8 {
	var sum uint32
	for y := 0; y < AEImageXSize; y++ {
		for x := 0; x < AEImageXSize; x++ {
			sum += uint32(img.GrayAt(x, y).Y)
		}
	}

	return uint8(sum / (AEImageXSize * AEImageYSize))
}
