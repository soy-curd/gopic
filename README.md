# go image processing for learning

## PGM(portable graymap format)の作成

```
brew install imagemagic
convert images/gopher.ppm -colorspace gray -compress none -scale 500x500 pgm:- > images/gopher.pgm # P2 pgm(テキスト形式)
convert images/gopher.ppm -colorspace gray -scale 500x500 pgm:- > images/gopher.pgm # P5 pgm(バイナリ形式)
```

## 参考サイト

- [碧色工房](https://www.mm2d.net/main/prog/c/image_io-01.html)

- [PGM 仕様](http://netpbm.sourceforge.net/doc/pgm.html)

## バグ？

### 旧review & ruby2.3
`画像を拡大するために、1画素を2x2=4画素にコピーして2倍のサイズにしてみます。`
という行がコンパイルできない。

```
! Package inputenc Error: Keyboard character used is undefined
(inputenc)                in inputencoding `utf8'.
```
-> utf8として認識できていない？？

### 新review & ruby2.3

```
! LaTeX Error: File `suffix.sty' not found.
```
https://tex.stackexchange.com/questions/111774/moderncv-latex-error-file-suffix-sty-not-found

-> これか？　（未対応）