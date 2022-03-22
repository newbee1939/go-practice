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
	// Go言語でポインタ変数を宣言するときは、型名の前に * を付ける必要がある
	// これにより、「ポインタ型」の変数として宣言できる
	var ptrA *string = &a
	fmt.Println(ptrA)

	// そのポインタ変数が指しているメモリの内容を参照する
	// 変数名そのものではなく、ポインタ変数を介して値を参照するので「間接参照」と呼ぶ
	// *は間接参照演算子と言う
	fmt.Println(*ptrA)

	// 間接参照を使って変数の値を書き換える
	*ptrA = "hoge"
	fmt.Println(*ptrA)

	// 構造体とは、いくつかの変数を1つにまとめて取り扱えるようにしたデータ型のこと
	// 構造体は、私たちプログラマが必要に応じて自由に定義することができるデータ型
	// 設計図のような役割
	// 構造体は、次のように struct キーワードを使って宣言する
	// 実際に構造体を使用するには、構造体をもとに変数を宣言する必要がある
	var person struct {
		name   string
		age    int
		gender string
		weight int
	}

	// フィールドに値をセットする（構造体の各フィールドに値を設定する）
	// これには二つの方法がある
	// 1つはドット演算子を使う方法
	// [構造体を格納した変数].[フィールド名]
	person.name = "newbe"
	person.age = 35
	person.gender = "female"
	person.weight = 34
	// %#v:Go言語のリテラル表現で対象データの情報を埋め込む
	fmt.Printf("%#v", person)

	// もう一つは構造体の作成時に初期化を行う方法
	hoge := struct {
		hige string
		huga int
	}{
		hige: "haa",
		huga: 23,
	}
	fmt.Printf("%#v", hoge)

	// 構造体にtypeを付ける -> 構造体を繰り返し使用することができる
	// type とは、既にある型（型リテラル）に対して新しい名前を与えるための機能
	// type [新しい名前] [既にある型]
	type Person struct {
		name   string
		age    int
		gender string
		weight int
	}
	// type を使用することで、これまで構造体を利用するときに struct {name string age int gender string weight int} と長ったらしく書いていた部分を、次のように Person という1つの単語に置き換えて書くことが可能に
	hanako := Person{
		name:   "Tanaka Hanako",
		age:    25,
		gender: "female",
		weight: 50,
	}
	println(hanako)
	// 順番通り値を定義した場合は省略することも可能
	yume := Person{"Yamada Yume", 20, "male", 60}
	println(yume)

	// メソッドについて
	// メソッドとは、特定の型に関連付けられた関数のこと

	// JavaやC++などのオブジェクト指向言語では、メソッドは「クラス」に関連付けられていて、クラスの内側でメソッドが定義されます
	// しかしGo言語の場合は、メソッドは構造体などの「型」に関連付けられ、型の外側でメソッドの定義が行われます
	// メソッドと型を関連付けるためにレシーバーというものを利用する
	// func (レシーバ値 レシーバ型) メソッド名(引数値 引数型) 戻り値の型

	// ここから！エラーを解消する！！
	hnako := Person{"Tanaka Hanako", 25, "female", 50}
	hnako.Greet() // => Hello!
}

// レシーバを渡す
// レシーバとは 型にメソッドを紐付ける役割を持った引数
// Greet メソッドは Person 型に紐付いている状態
func (p Person) Greet() {
	fmt.Println("Hello!")
}
