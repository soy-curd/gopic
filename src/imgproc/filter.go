package imgproc

import (
	"./util"
)

// Laplacian applys edge filter to image
func (img *Pgm) Laplacian() {
	laplacianFilter := [][]int{{1, 1, 1}, {1, -8, 1}, {1, 1, 1}}
	img.data = applyFilter(img, laplacianFilter, img.tone)
}

// PatternHorizontal applys horizontal filter to image
func (img *Pgm) PatternHorizontal() {
	horizontalFilter := [][]int{{-1, -1, -1}, {2, 2, 2}, {-1, -1, -1}}
	// varticalFilter := [][]int{{-1, 2, -1}, {-1, 2, -1}, {-1, 2, -1}}
	img.data = applyFilter(img, horizontalFilter, img.tone)
}

func applyFilter(img *Pgm, filter [][]int, tone int) [][]byte {
	buf := [][]byte{}
	buf = append(buf, make([]byte, len(img.data[0])))

	// ignore periphery of image
	for i := 1; i < img.height-1; i++ {
		buf = append(buf, []byte{})
		buf[i] = append(buf[i], byte(0))
		for j := 1; j < img.width-1; j++ {
			buf[i] = append(buf[i], byte(calcFilter(img.data, i, j, filter, img.tone)))
		}
		buf[i] = append(buf[i], byte(0))
	}
	buf = append(buf, make([]byte, len(img.data[0])))
	return buf
}

func calcFilter(data [][]byte, i int, j int, filter [][]int, tone int) int {
	p := 0
	buf := []int{}

	for x := 0; x < len(filter); x++ {
		for y := 0; y < len(filter[0]); y++ {
			p += int(data[i+(x-1)][j+(y-1)]) * filter[x][y]
			buf = append(buf, int(data[i+(x-1)][j+(y-1)]))
		}
	}
	p = int(float64(p) / float64((len(filter) * len(filter[0]))))
	return util.Min(util.Max(p, 0), tone)
}
