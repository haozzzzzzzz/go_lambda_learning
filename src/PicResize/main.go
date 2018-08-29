package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	file, err := os.Open("./src/PicResize/shuangji.png")
	if nil != err {
		log.Fatal(err)
		return
	}

	img, str, err := image.Decode(file)
	file.Close()
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println(str)

	m := resize.Resize(100, 0, img, resize.Lanczos3)
	out, err := os.Create("./src/PicResize/new_shuangji.png")
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
