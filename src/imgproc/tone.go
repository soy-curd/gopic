package imgproc

import (
	"./util"
)

// CorrectTone changes image to high contrast
func (img *Pgm) CorrectTone() {
	max := 0
	min := img.tone

	// check max/min tone in image
	for i := 0; i < img.height; i++ {
		for j := 0; j < img.width; j++ {
			max = util.Max(int(img.data[i][j]), max)
			min = util.Min(int(img.data[i][j]), min)
		}
	}

	for i := 0; i < img.height; i++ {
		for j := 0; j < img.width; j++ {
			img.data[i][j] = byte(
				normalize(
					int(img.data[i][j]), min, max, img.tone))
		}
	}
}

func normalize(p int, min int, max int, toneMax int) int {
	tone := (float64(p-min) / float64(max-min)) * float64(toneMax)
	return int(tone)
}
