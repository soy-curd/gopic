package main

import (
	"io/ioutil"

	"./imgproc"
)

func main() {
	// imgproc.Dump("images/gopher.pgm")
	// data, err := ioutil.ReadFile("images/gopher.pgm")
	imgproc.Dump("images/lena.pgm")
	data, err := ioutil.ReadFile("images/lena.pgm")
	// data, err := ioutil.ReadFile("images/gopher_low_contrast.pgm")
	if err != nil {
		panic(err)
	}

	image := imgproc.DecodePgm(data)
	// fmt.Println(image)
	// image.Brighten(100)
	// image.Flip()
	// image.Reverse()
	// image.Enlarge()
	// image.CorrectTone()
	// image.Binarization(100)
	image.DiffuseError()
	out := imgproc.EncodePgm(image)
	ioutil.WriteFile("images/gopher_out.pgm", out, 0666)
	imgproc.Dump("images/gopher_out.pgm")
	// fmt.Println(image)
}
