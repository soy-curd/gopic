
= 画像ファイルを読み込んでみよう

== ダンプする


Goを用いると、簡単に画像のダンプを見ることができます。@<href>{https://github.com/soy-curd/gopic/blob/master/images/gopher.pgm,こちら}からダウンロードした画像の、中身を見ていきましょう。


//emlist[][src/imgproc/pgm.go]{
data, _ := ioutil.ReadFile(filepath)
fmt.Printf("%s", hex.Dump(data[:100]))
fmt.Println("file dump finished!")
//}


こちらを実行すると、


//emlist{
00000000  50 35 0a 23 20 38 2d 62  69 74 20 70 70 6d 20 52  |P5.# 8-bit ppm R|
00000010  47 42 0a 35 30 30 20 35  30 30 0a 32 35 35 0a ff  |GB.500 500.255..|
00000020  ff ff ff ff ff ff ff ff  ff ff ff ff ff ff ff ff  |................|
....|
//}


のようなダンプが出力されます。



この画像はPGM（portable gray map）ファイルフォーマットで、ヘッダー部の後に画像の画素の階調値のデータが続く構造となっています。



ヘッダー部は上記の画像の場合、


//emlist{
P5
# 8-bit ppm RGB
500 500
255
//}


の部分で、それぞれ PGM の形式マジックナンバー（P5 は白黒画像のバイナリ形式）、コメント、幅・高さのサイズ、階調値の最大値、となります。



続く画素のデータについては、ダンプを見ると ff ff ff ...と連続しているのがわかります。PGM のバイナリ形式は 1byteを1画素とするので、ffは255を示します。階調数は明るさの段階を示し、画素は左上から右下に向かって並んでいるので、この画像の左上端は最も明るい白ということになります。（実際にpgmのフォーマットが開ける画像ビューワで確認してみてください）


== 画像をパースする


画像処理を行うためには、プログラムで取扱いやすい形で画像を読み込む必要があります。



画像は1ファイルを、以下の構造体に格納することとします。


//emlist{
type Pgm struct {
    width  int
    height int
    size   int  // width * height
    tone   int
    data   [][]byte
}
//}


ここで、toneは階調の最大値を示します。dataはbyte型の２次元スライスとして格納します。



次に、画像をパースしていきましょう。


//emlist[][golang]{

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

//}

//emlist[][golang]{
func main() {
    data, err := ioutil.ReadFile("images/gopher.pgm")
    if err != nil {
        panic(err)
    }

    image := imgproc.DecodePgm(data)
}
//}


ここで、getNextToken関数は読み込んだデータから最初の文字列を取り出します。この関数の中ではgetNextNonSpaceChar関数を用いて、コメントを読み飛ばしています。getNextInt関数では、getNextToken関数で得た文字列を数値に変換しています。そして、公開関数であるDecodePgm関数では、Pgmのフォーマットに従ってデータを読み込み、先程定義したPgm構造体に格納しています。



※ pgm読み込み処理は、@<href>{https://www.mm2d.net/main/prog/c/image_io-01.html,碧色工房}を参考にしました。


== 画像を保存する


画像を保存する場合は、Pgm構造体に格納したデータをbyte列に戻すだけで良いです。


//emlist{
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
//}

//emlist{
    out := imgproc.EncodePgm(image)
    ioutil.WriteFile("images/gopher_out.pgm", out, 0666)
//}


以上が、pgmデータの読み込みと書き出しの処理になります。次の章では、実際に読み出した画像を様々な方法で変換していきます。

