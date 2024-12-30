/**
静的ファイルを返却するサーバー

実行方法
- ターミナル1で "go run <ファイル名>" を実行してサーバーを立ち上げる
- Webブラウザで "http://localhost:8080/hello.html" にアクセス
*/

package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}