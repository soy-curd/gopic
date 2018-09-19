package main

import (
	"io/ioutil"

	"./imgproc"
)

func main() {
	imgproc.Dump("images/gopher.pgm")
	data, err := ioutil.ReadFile("images/gopher.pgm")
	if err != nil {
		panic(err)
	}

	image := imgproc.DecodePgm(data)
	// image.Brighten(100)
	image.Flip()
	out := imgproc.EncodePgm(image)
	ioutil.WriteFile("images/gopher_out.pgm", out, 0666)
	// fmt.Println(image)
}
