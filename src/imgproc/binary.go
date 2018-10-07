package imgproc

import "./util"

// Binarization converts image to binary value
func (img *Pgm) Binarization(threshold int) {
	for i := 0; i < img.height; i++ {
		for j := 0; j < img.width; j++ {
			if int(img.data[i][j]) > threshold {
				img.data[i][j] = byte(img.tone)
			} else {
				img.data[i][j] = byte(0)
			}
		}
	}
}

// DiffuseError converts image to binary value by Error Diffusion Method
func (img *Pgm) DiffuseError() {
	half := img.tone / 2
	for i := 0; i < img.height; i++ {
		for j := 0; j < img.width; j++ {
			applyError(img.data, i, j, img.tone)
			p := int(img.data[i][j])
			if p > half {
				img.data[i][j] = byte(img.tone)
			} else {
				img.data[i][j] = byte(0)
			}
		}
	}
}

func applyError(data [][]byte, i int, j int, tone int) {
	p := int(data[i][j])
	var error int
	if p < tone-p {
		error = p
	} else {
		error = p - tone
	}

	if j+1 < len(data[i]) {
		data[i][j+1] = byte(util.Min(
			int(0.3*float64(error))+int(data[i][j+1]),
			tone))
	}
	if i+1 < len(data) && j > 0 {
		data[i+1][j-1] = byte(util.Min(
			int(0.2*float64(error))+int(data[i+1][j-1]),
			tone))
	}
	if i+1 < len(data) {
		data[i+1][j] = byte(util.Min(
			int(0.3*float64(error))+int(data[i+1][j]),
			tone))
	}
	if i+1 < len(data) && j+1 < len(data[i]) {
		data[i+1][j+1] = byte(util.Min(
			int(0.3*float64(error))+int(data[i+1][j+1]),
			tone))
	}
}
