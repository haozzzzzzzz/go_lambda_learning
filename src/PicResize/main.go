package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	file, err := os.Open("./src/PicResize/HappyFace.jpg")
	if nil != err {
		log.Fatal(err)
		return
	}

	img, err := jpeg.Decode(file)
	file.Close()
	if nil != err {
		log.Fatal(err)
		return
	}

	m := resize.Resize(100, 0, img, resize.Lanczos3)
	out, err := os.Create("./src/PicResize/new_HappyFace.jpg")
	if nil != err {
		log.Fatal(err)
		return
	}
	defer out.Close()

	err = jpeg.Encode(out, m, nil)
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println("success")
}
