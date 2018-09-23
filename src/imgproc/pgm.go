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

// Dump prints hexdump
func Dump(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", hex.Dump(data[:100]))
	fmt.Println("file dump finished!")
	return data
}

// DecodePgm gets image structure
func DecodePgm(data []byte) Pgm {
	img := Pgm{}

	buf := bytes.NewReader(data)
	getNextToken(buf) // not used
	img.width = getNextInt(buf)
	img.height = getNextInt(buf)
	img.size = img.width * img.height
	img.tone = getNextInt(buf)
	img.data = [][]byte{}
	imageBytes := make([]byte, buf.Len())
	buf.Read(imageBytes)

	for i := 0; i < img.height; i++ {
		img.data = append(img.data, imageBytes[i*img.width:(i+1)*img.width])
	}
	return img
}

// EncodePgm makes pgm img
func EncodePgm(img Pgm) []byte {
	var buf bytes.Buffer
	buf.Write([]byte("P5\n"))
	buf.Write([]byte(strconv.Itoa(img.width) + " "))
	buf.Write([]byte(strconv.Itoa(img.height) + "\n"))
	buf.Write([]byte(strconv.Itoa(img.tone) + "\n"))
	for i := 0; i < img.height; i++ {
		buf.Write(img.data[i])
	}
	return buf.Bytes()
}

func getNextNonSpaceChar(reader *bytes.Reader) byte {
	var b byte
	var err error
	for {
		b, err = reader.ReadByte()
		if err != nil {
			break
		}
		// skip comment line
		if b == '#' {
			skipUntilLineReturn(reader)
			continue
		}

		if !unicode.IsSpace(rune(b)) {
			break
		}
	}
	return b
}

func skipUntilLineReturn(reader *bytes.Reader) {
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		if b == '\n' {
			break
		}
	}
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
