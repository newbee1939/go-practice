package main

import "fmt"

func main() {
	var a int
	// &はアドレス演算子。値のアドレスを示す
	fmt.Scan(&a)
	// 3を出力
	fmt.Println(a)
	// 0xc0000aa000を出力
	fmt.Println(&a)
}
