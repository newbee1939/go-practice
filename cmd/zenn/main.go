package main

import "fmt"

func main() {
	// とってもやさしいGo言語入門:https://zenn.dev/ak/articles/1fb628d82ed79b

	// 変数のメモリアドレスを表示する -> &を使う
	var a string = "test"
	var num int = 1
	fmt.Println(&a)
	fmt.Println(&num)

	// 取得したメモリアドレスを変数に格納する
	// ptrA のように「メモリアドレスを入れることのできる変数」のことを「ポインタ変数」と呼ぶ
	// Go言語におけるポインタはある特定のメモリアドレスを表す
	// ptrA は a のメモリアドレスをポイントする（指し示す）役割を持っている
	var ptrA *string = &a
	fmt.Println(ptrA)

	// ポインタ変数から値を取り出す
	// *は間接参照演算子と言う
	fmt.Println(*ptrA)
}
