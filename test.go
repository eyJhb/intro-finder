package main

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/Nr90/imgsim"
)

func worker(id int, jobs <-chan int, results chan<- imgsim.Hash) {
	for j := range jobs {
		if j%100 == 0 {
			fmt.Printf("Doing job: %d\n", j)
		}
		imgfile, err := os.Open(fmt.Sprintf("images5/out%d.jpeg", j))
		// defer imgfile.Close()
		if err != nil {
			panic(err)
		}
		img, err := jpeg.Decode(imgfile)
		if err != nil {
			panic(err)
		}
		ahash := imgsim.AverageHash(img)
		// ahash := imgsim.DifferenceHash(img)
		results <- ahash
		imgfile.Close()
	}
}

func main() {
	imgfile, err := os.Open(fmt.Sprintf("out0020.jpeg"))
	// defer imgfile.Close()
	if err != nil {
		panic(err)
	}
	img, err := jpeg.Decode(imgfile)
	if err != nil {
		panic(err)
	}
	ahash := imgsim.AverageHash(img)
	// ahash := imgsim.DifferenceHash(img)
	fmt.Println(ahash)
	imgfile.Close()
	// setup our jobs and results
	jobs := make(chan int, 11000)
	results := make(chan imgsim.Hash, 11000)

	// make our workers
	for w := 0; w < 8; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 10790; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= 10790; r++ {
		<-results
	}
}
