/**
文字列を返却するサーバー

実行方法
- ターミナル1で "go run <ファイル名>" を実行してサーバーを立ち上げる
- ターミナル2で "curl http://localhost:8080" を実行してリクエスト
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintは第一引数に指定された出力先に任意の文字列を書き込む (fmt.Printの上位互換)
	// http.ResponseWriter(HTTPレスポンス)に文字列を書き込む
	fmt.Fprint(w, "Hello, Web application!")
}

func main() {
	// HTTPリクエストに対して実行する処理を定義
	// "/" にアクセスが来た場合は hello関数を実行する
	http.HandleFunc("/", hello)

	// HTTPサーバーを起動
	// 8080ポートで待ち受ける
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}