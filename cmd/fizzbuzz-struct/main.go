package main

// Goでは構造体がフィールドとメソッドを定義できる。オブジェクト指向のクラスと部分的に似たような使い方ができる
// 構造体がまだ掴めない。。

func main() {
	// 構造体を生成して初期化
	fb := FizzBuzz{1, 100}
	// 実行
	fb.Run()
}

// FizzBuzz構造体を定義
type FizzBuzz struct {
	Cur int
	Max int
}

// 構造体FizzBuzzで処理実行のメソッド
// *の意味とかちゃんと理解する
func (p *FizzBuzz) Run() {
	// ここから！！
}
