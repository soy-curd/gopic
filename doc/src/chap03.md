# 階調補正と二値化

## 階調補正

![薄暗いGopherくん](./images/gopher_low_contrast.png)

例えば明暗のはっきりしない画像があったとき、これをくっきりとした画像にするにはどうしたら良いでしょうか。

例えば上の Gopher くんは階調の分布が中央付近に集中しています。階調の最大値は 211、最小値は 160 です。この階調を、0~255 の間に再分布することによって、コントラストを明瞭にすることができます。

```golang
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
```

ここで normalize 関数は、((対象の画素の階調値 - 画像の階調の最小値) / (画像の階調の最大値 - 画像の階調の最小値)) \* pgm の階調の最大値(255)という計算をしています。

## 二値化

今まで扱ってきた画像は画素の明るさに階調がありましたが、例えば印刷用の画像が必要な場合など、階調が 0 か 255 しかない二値化画像のほうが扱い安い場合があります。

まずは単純に、ある閾値を超えたら階調値を 255、超えなかったら 0 にする関数を作ってみましょう。

```golang
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
```

![二値化Gopherくん](./images/gopher_binary.png)

しかし、この場合、適切な閾値とはどのような値でしょうか。もともとの画像が全体的に明るい場合や暗い場合に、アドホックに閾値を選ぶ必要がでてきてしまいます。また、もともと 0 から 255 の間にあった中間的な階調を表現することができません。この問題の対応として、疑似階調表現という手法があります。疑似階調表現では、画素を n×n の領域に区切り、その領域で 0 ないしは 255 の階調を持った画素を配置していきます。

疑似階調表現の手法の一つの誤差拡散法では、注目する画素の階調と 0 ないしは 255 との誤差（より小さい値）を隣接する画素に加えていきます。

```golang
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
```

上の処理では、applyError 関数で誤差を計算し、隣接する画素に加算しています。誤差は 255 に近い場合は生の値、0 に近い場合は負の値になります。隣接する画素へは一定の係数をかけた値を加算していて、それぞれ、右隣は 0.3、左下は 0.2、下は 0.3、右下は 0.2 としています。

![誤差拡散Gopherくん](./images/gopher_diff.png)

もともとの輪郭以外の要素も、二値化画像として表現されていることがわかります。波模様として出現しているのはモアレと呼ばれるもので、走査方向を調整することなどによって低減することができます。
