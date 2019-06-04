package hash

import (
	"fmt"
	"image"

	"github.com/nfnt/resize"
)

const (
	DEImageXSize = 17
	DEImageYSize = 16
)

func DEHash(img image.Image) [4]uint64 {
	// resize our image
	resImage := resize.Resize(DEImageXSize, DEImageYSize, img, resize.NearestNeighbor)

	// make grayscale
	grayImage := rgbaToGray(resImage)

	// loop over our image, and set hash accordingly
	hash := [5]uint64{}
	var currentIndex uint8
	var left, right uint8

	for y := 0; y < DEImageYSize; y++ {
		// set our initial left value
		left = grayImage.GrayAt(0, y).Y
		for x := 1; x < DEImageXSize; x++ {
			right = grayImage.GrayAt(x, y).Y
			if right < left {
				hash[currentIndex] |= 1
				// fmt.Printf("1")
			} else {
				// fmt.Printf("0")
			}
			hash[currentIndex] <<= 1

			if ((y*DEImageYSize)+x)%64 == 0 {
				fmt.Println("changed..")
				currentIndex++
			}
		}
	}
	fmt.Println()
	fmt.Println(hash)
	fmt.Printf("%064b%064b%064b%064b\n", hash[0], hash[1], hash[2], hash[3])

	// fmt.Printf("%064b\n", hash)

	return [4]uint64{hash[0], hash[1], hash[2], hash[3]}
	// return hash
}

func DiffE(hash1, hash2 [4]uint64) uint8 {
	// holds the distance between the two
	var distance uint8

	// init our mask, and set initial value of 1

	var mask uint64
	for j := 0; j < 4; j++ {
		mask = 1
		for i := 0; i < 64; i++ {
			// mask our hash1 and hash2, see if they are different
			// if se increment the distance
			if hash1[j]&mask != hash2[j]&mask {
				distance++
			}

			// move our mask the the left
			mask <<= 1
		}
	}

	return distance
}
