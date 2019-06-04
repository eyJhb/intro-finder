package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"intro-finder/hash"
	"os"
)

// -rw-r--r-- 1 eyjhb users  7624 Apr 24 13:18 60-0001.jpg
// -rw-r--r-- 1 eyjhb users 10764 Apr 24 13:26 60-0003.jpg
// -rw-r--r-- 1 eyjhb users 13016 Apr 24 13:10 80-0002.jpg
// -rw-r--r-- 1 eyjhb users 10288 Apr 24 13:11 80-0003.jpg
// -rw-r--r-- 1 eyjhb users 12325 Apr 24 13:16 80-0004.jpg

func main() {
	// t5 := hash.AEHash(readFile("data/p42/60-0001.jpg"))
	// t6 := hash.AEHash(readFile("data/p42/60-0003.jpg"))
	// fmt.Println(hash.DiffE(t5, t6))
	// return

	t1 := hash.DEHash(readFile("data/p42/60-0001.jpg"))
	t2 := hash.DEHash(readFile("data/p42/60-0003.jpg"))
	f1 := hash.DEHash(readFile("data/p42/80-0002.jpg"))
	f2 := hash.DEHash(readFile("data/p42/80-0003.jpg"))
	f3 := hash.DEHash(readFile("data/p42/80-0004.jpg"))
	h1 := hash.DEHash(readFile("data/p42/90-0001.jpg"))
	h2 := hash.DEHash(readFile("data/p42/90-0002.jpg"))
	v1 := hash.DEHash(readFile("data/p42/vige-0001.jpg"))
	s1 := hash.DEHash(readFile("data/p42/stop-0002.jpg"))
	// fmt.Printf("%064b\n", t1)
	// fmt.Printf("%064b\n", t2)
	// fmt.Printf("%064b\n", f1)
	// fmt.Printf("%064b\n", f2)
	// fmt.Printf("%064b\n", f3)
	// fmt.Printf("%064b\n", h1)
	// fmt.Printf("%064b\n", h2)
	// fmt.Printf("%064b\n", v1)
	// fmt.Printf("%064b\n", s1)

	fmt.Println(hash.DiffE(t1, t2))
	fmt.Println("--")
	fmt.Println(hash.DiffE(t1, f1))
	fmt.Println(hash.DiffE(t1, f2))
	fmt.Println(hash.DiffE(t1, f3))
	fmt.Println("--")
	fmt.Println(hash.DiffE(t2, f1))
	fmt.Println(hash.DiffE(t2, f2))
	fmt.Println(hash.DiffE(t2, f3))
	fmt.Println("--")
	fmt.Println(hash.DiffE(f1, f2))
	fmt.Println(hash.DiffE(f1, f3))
	fmt.Println(hash.DiffE(f2, f3))
	fmt.Println("--")
	fmt.Println(hash.DiffE(h1, h2))
	fmt.Println(hash.DiffE(h1, f1))
	fmt.Println(hash.DiffE(h1, f2))
	fmt.Println(hash.DiffE(h1, f3))
	fmt.Println(hash.DiffE(h2, f1))
	fmt.Println(hash.DiffE(h2, f2))
	fmt.Println(hash.DiffE(h2, f3))
	fmt.Println("--")
	fmt.Println(hash.DiffE(t1, v1))
	fmt.Println(hash.DiffE(v1, s1))

	// fmt.Println(hash.Diff(h1, h2))
	// fmt.Println(hash.Diff(h1, f1))
	// fmt.Println(hash.Diff(h1, f2))
	// fmt.Println(hash.Diff(h1, f3))
	// fmt.Println(hash.Diff(h2, f1))
	// fmt.Println(hash.Diff(h2, f2))
	// fmt.Println(hash.Diff(h2, f3))
	// fmt.Println(hash.Diff(h1, v1))
	// fmt.Println(hash.Diff(h1, s1))

	// ahash := hash.AHash(readFile(filename))

}

func readFile(filename string) image.Image {
	infile, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer infile.Close()

	// Must specifically use jpeg.Decode() or it
	// would encounter unknown format error
	src, _, err := image.Decode(infile)
	if err != nil {
		panic(err.Error())
		return nil
	}

	return src
}
