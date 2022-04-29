package main

import "fmt"

func main() {
	// Goの基礎文法を世界一分かりやすくまとめてみた
	// とってもやさしいGo言語入門:https://zenn.dev/ak/articles/1fb628d82ed79b

	// 変数のメモリアドレスを表示する -> &を使う
	// メモリアドレスは16進数で表現される
	var a string = "test"
	var num int = 1
	fmt.Println(&a)
	fmt.Println(&num)

	// 取得したメモリアドレスを変数に格納する
	// ptrA のように「メモリアドレスを入れることのできる変数」のことを「ポインタ変数」と呼ぶ
	// Go言語における「ポインタ」はある特定の「メモリアドレス」を表す
	// ptrA は a のメモリアドレスをポイントする（指し示す）役割を持っている
	// Go言語でポインタ変数を宣言するときは、型名の前に * を付ける必要がある
	// これにより「ポインタ型」の変数として変数を宣言できる
	var ptrA *string = &a
	fmt.Println(ptrA)

	// 間接参照
	// そのポインタ変数が指しているメモリの内容を参照する
	// 変数名そのものではなく、ポインタ変数を介して値を参照するので「間接参照」と呼ぶ
	// *は間接参照演算子と言う
	fmt.Println(*ptrA) // ptrAに格納したアドレスが指し示す値を出力することができる

	// 間接参照を使って変数の値を書き換えることもできる
	*ptrA = "hoge"
	fmt.Println(*ptrA)

	// 構造体とは、いくつかの変数を1つにまとめて取り扱えるようにした（関連するデータを一つにまとめて宣言できる）データ型のこと
	// 構造体は、私たちプログラマが必要に応じて自由に定義することができるデータ型
	// 設計図のような役割
	// 構造体は、次のように struct キーワードを使って宣言する
	// 実際に構造体を使用するには、構造体をもとに変数を宣言する必要がある
	// name が string 型で age が int 型で gender が string 型で weight が int 型のフィールドを持った構造体の型を person という変数名で宣言する
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

	// もう一つは構造体の作成時に「初期化」を行う方法
	hoge := struct {
		hige string
		huga int
	}{
		hige: "haa",
		huga: 23,
	}
	fmt.Printf("%#v", hoge)

	// 構造体にtypeを付ける -> 構造体を繰り返し使用することができる（基本この方法を使う）
	// type とは、既にある型（型リテラル）に対して新しい名前を与えるための機能（intとかstringみたいな新たな型を定義するイメージかな）
	// type [新しい名前] [既にある型]
	type Person struct {
		name   string
		age    int
		gender string
		weight int
	}
	// type を使用することで、これまで構造体を利用するときに
	// struct {name string age int gender string weight int} と長ったらしく書いていた部分を、
	// 次のように Person という1つの単語に置き換えて書くことが可能に
	hanako := Person{
		name:   "Tanaka Hanako",
		age:    25,
		gender: "female",
		weight: 50,
	}
	println(hanako)
	// 順番通り値を定義した場合はフィールド名を省略することも可能
	yume := Person{"Yamada Yume", 20, "male", 60}
	println(yume)

	// メソッドについて
	// メソッドとは、「特定の型に関連付けられた関数」のこと

	// JavaやC++などのオブジェクト指向言語では、メソッドは「クラス」に関連付けられていて、クラスの内側で「メソッド」が定義されます
	// しかしGo言語の場合は、メソッドは構造体などの「型」に関連付けられ、型の外側で「メソッド」の定義が行われます（この辺りオブジェクト指向とは違うなあ。。）
	// メソッドと型を関連付けるためにレシーバーというものを利用する
	// func (レシーバ値 レシーバ型) メソッド名(引数値 引数型) 戻り値の型

	hnako := Person{"Tanaka Hanako", 25, "female", 50}
	hnako.Greet() // => Hello!

	// レシーバを渡す
	// レシーバとは「型にメソッドを紐付ける役割を持った引数」
	// これでGreet メソッドは Person 型に紐付いている状態になる
	// pは、このメソッドを呼び出している変数のデータ（hnako）をそのまま受け取ることができる
	func (p Person) Greet() {
		fmt.Println("Hello!")
	}

	// ポンターレシーバ
	// フィールドの値を「直接操作」するメソッドについて
	func (p Person) Eat() {
		p.weight += 1
	}

	hanako.Eat() // Eatメソッドの実行
 	fmt.Println("食後の体重は", hanako.weight) // => 食後の体重は 50 <- 増えてない
	// レシーバの型は、「値型」か「ポインタ型」かどうかで挙動が変わってくる
	// 値型 = メソッドの呼び出し時に「レシーバそのもののコピー」が発生する（メソッド内でフィールドの値を変更してもその変更はコピーに対して行われるため、元の実体の値は変わらない）
	// ポインタ型 = ポインタのレシーバを受け取るため、メソッド内の「実体」に対して変更処理を書くことができる
	// つまり、メソッド内で実体の変更をしたければ、レシーバにはポインタ型の変数を渡さないといけないということ
	// 以下の例なら書き換えが可能
	func (p *Person) Eat() {
		(*p).weight += 1
	}
	hanako.Eat() // Eatメソッドの実行
 	fmt.Println("食後の体重は", hanako.weight) //メソッドの呼び出し元のコードは何も変更する必要がない点に注意！！
	// たとえ呼び出し元の変数がポインタ型じゃなくても、レシーバがポインタ型である場合は Go言語が勝手に判断して、それをポインタ型として解釈してくれる

	// インターフェースについて
	// インターフェース = 型が持つべきメソッドを規約として定めたもの
	// Goの場合、ある特定の型がどのようなメソッドを持つべきかを定めるためにインターフェースを使う

	// インターフェースの定義方法
	// interfaceキーワードを使う
	type 型名 interface {
		メソッド名(引数名 引数の型) 戻り値の型
	}
	// Vehicle型の場合
	type Vehicle interface {
		Accelerate()
		Brake()
	}
}
