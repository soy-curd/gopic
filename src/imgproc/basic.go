package imgproc

import "./util"

// Brighten converts image brightness
func (img *Pgm) Brighten(shift int) {
	for i := 0; i < img.height; i++ {
		for j := 0; j < img.width; j++ {
			img.data[i][j] = byte(util.Min(int(img.data[i][j])+shift, img.tone))
		}
	}
}

// Flip flips image horizontal
func (img *Pgm) Flip() {
	var fliped [][]byte

	for i := 0; i < img.height; i++ {
		fliped = append(fliped, []byte{})
		for j := 0; j < img.width; j++ {
			fliped[i] = append(fliped[i], img.data[i][img.width-j-1])
		}
	}
	img.data = fliped
}

// Reverse reverses image tone
func (img *Pgm) Reverse() {
	for i := 0; i < img.height; i++ {
		for j := 0; j < img.width; j++ {
			img.data[i][j] = byte(img.tone - int(img.data[i][j]))
		}
	}
}

// Enlarge converts image size double
func (img *Pgm) Enlarge() {
	var resized [][]byte
	var scale = 2

	for i := 0; i < img.height*scale; i++ {
		resized = append(resized, []byte{})
		for j := 0; j < img.width*scale; j++ {
			resized[i] = append(resized[i], img.data[i/scale][j/scale])
		}
	}
	img.width = img.width * scale
	img.height = img.height * scale
	img.size = img.width * img.height
	img.data = resized
}
