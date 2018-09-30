# はじめに

Gopherくんというキャラクターを知っているでしょうか？

![Gopherくん](./images/gopher.png)

Gopherくんはプログラミング言語Goのマスコットキャラクターです。この愛嬌のある表情を見ると、このGopherくんを画像処理してみたくなりますね。せっかくですから、Goを使ってGopherくんを処理していきましょう。

GoはGoogleで開発されたプログラミング言語です。Goはマシン語にコンパイルされてから実行されるため、PythonやRubyなどのインタプリタ言語と比較して高速です。また、ガベージコレクションや型推論等の仕組みを持っているため、C言語よりも手軽にプログラミングすることができます。Go自体は[Tour of Go](https://go-tour-jp.appspot.com/welcome/1)で学ぶことができますので、Goに触れたことがない方はまずそちらを一通り触ってみると良いかと思います。

## （画像を扱う前に）Goでバイナリデータを扱う

画像処理ではバイナリデータを扱うことが必要です。画像ファイルを扱う前に、簡単なバイナリ処理をしてみましょう。

Goでのbyte型は8ビットunsigned intのエイリアスです。例えば例えばバイト列を、次のように初期化できます。

```
a := []byte{0, 255}  // 10進整数で値指定
b := []byte{0x61, 0x62}  // 16進数で値指定
```

8ビットunsigned intの範囲から外れた値で初期化しようとするとコンパイルエラーになります。

```
a := []byte{256}
b := []byte{-1}
```

printする場合は、フォーマット指定子に`%x`や`%q`を用いるか、"encoding/hex"パッケージの dump 関数を用いると便利です。

```
a := []byte{0, 255}  // 10進整数で値指定
b := []byte{0x61, 0x62}  // 16進数で値指定
fmt.Printf("%x, %x\n", a, b)  // 00ff, 6162
fmt.Printf("%q, %q\n", a, b)  // "\x00\xff", "ab"
print(hex.Dump([]byte("helloこんにちは")))
00000000  68 65 6c 6c 6f e3 81 93  e3 82 93 e3 81 ab e3 81  |hello...........|
00000010  a1 e3 81 af                                       |....|
```

Goでは**bytes**パッケージにbyte列を操作する関数が提供されています。
例えば、bytes.Equal関数を用いるとbyte列を比較できます。

```
c := []byte{0, 1, 2, 3}
d := []byte{0, 1, 2, 3}
print(bytes.Equal(c, d))  // true
```

また、bytesパッケージにはbytes.Buffer型という読み書き可能な可変サイズバッファが定義されています。bytes.ReadByte関数を用いると、バッファから1バイト読み進めることができます。

```
e := bytes.NewBuffer([]byte("helloこんにちは"))
next, _ := e.ReadByte()
fmt.Printf("%q, %d\n", next, e.Len())
next, _ = e.ReadByte()
fmt.Printf("%q, %d\n", next, e.Len())
next, _ = e.ReadByte()
fmt.Printf("%q, %d\n", next, e.Len())
// "h", 19
// "e", 18
// "l", 17
```


それでは次の章から、実際に画像ファイルを扱ってみましょう。
