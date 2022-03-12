package main

import (
	"fmt"
	// I/O(入出力)関連の便利なパッケージを提供するパッケージ
	"io/ioutil"
	"net/http"
)

func main() {
	// アクセスするURL
	p := "https://engilaboo.com"
	// 指定のアドレスにアクセスする
	re, er := http.Get(p)
	if er != nil {
		// プログラムを強制終了する
		panic(er)
	}
	// deferでは「プログラムを終了する際に必ず実行する処理」を指定できる
	// ReadCloserの解放処理
	defer re.Body.Close()

	// Getで得られたResponseからコンテンツを取り出す
	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}

	// 出力
	fmt.Println(string(s))
}
