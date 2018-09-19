package imgproc

// Brighten convert image brightness
func (image *Pgm) Brighten(shift int) {
	for i := 0; i < image.height; i++ {
		for j := 0; j < image.width; j++ {
			image.data[i][j] = byte(max(int(image.data[i][j])+shift, image.tone))
		}
	}
}

// Flip horizontal
func (image *Pgm) Flip() {
	var fliped [][]byte

	for i := 0; i < image.height; i++ {
		fliped = append(fliped, []byte{})
		for j := 0; j < image.width; j++ {
			fliped[i] = append(fliped[i], image.data[i][image.width-j-1])
		}
	}
	image.data = fliped
}

func max(a, b int) int {
	if a < b {
		return a
	}
	return b
}
