package imgproc

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"unicode"
)

// Pgm is pgm data structure
type Pgm struct {
	width  int
	height int
	size   int // width * height
	tone   int
	data   [][]byte
}

// Dump print hexdump
func Dump(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", hex.Dump(data[:100]))
	fmt.Println("file dump finished!")
	return data
}

// DecodePgm get image structure
func DecodePgm(data []byte) Pgm {
	image := Pgm{}

	buf := bytes.NewReader(data)
	getNextToken(buf) // not used
	image.width = getNextInt(buf)
	image.height = getNextInt(buf)
	image.size = image.width * image.height
	image.tone = getNextInt(buf)
	image.data = [][]byte{}
	imageBytes := make([]byte, buf.Len())
	buf.Read(imageBytes)

	for i := 0; i < image.height; i++ {
		image.data = append(image.data, imageBytes[i*image.width:(i+1)*image.width])
	}
	return image
}

// EncodePgm make pgm image
func EncodePgm(image Pgm) []byte {
	var buf bytes.Buffer
	buf.Write([]byte("P5\n"))
	buf.Write([]byte(strconv.Itoa(image.width) + " "))
	buf.Write([]byte(strconv.Itoa(image.height) + "\n"))
	buf.Write([]byte(strconv.Itoa(image.tone) + "\n"))
	for i := 0; i < image.height; i++ {
		buf.Write(image.data[i])
	}
	return buf.Bytes()
}

func getNextNonSpaceChar(buf *bytes.Reader) byte {
	comment := false
	var b byte
	var err error
	for {
		b, err = buf.ReadByte()
		if err != nil {
			break
		}
		if comment {
			if b == '\n' {
				comment = false
			}
			continue
		}
		if b == '#' {
			comment = true
			continue
		}

		if !unicode.IsSpace(rune(b)) {
			break
		}
	}
	return b
}

func getNextToken(buf *bytes.Reader) string {
	var b byte
	var err error
	var tokenBuf bytes.Buffer
	b = getNextNonSpaceChar(buf)
	tokenBuf.WriteByte(b)

	for {
		b, err = buf.ReadByte()
		if err != nil {
			break
		}
		if unicode.IsSpace(rune(b)) {
			break
		}
		tokenBuf.WriteByte(b)
	}
	token := tokenBuf.String()
	return token
}

func getNextInt(buf *bytes.Reader) int {
	token := getNextToken(buf)
	intValue, err := strconv.Atoi(token)
	if err != nil {
		panic(err)
	}
	return intValue
}
