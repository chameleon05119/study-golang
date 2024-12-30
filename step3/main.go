/**
ToDoアプリケーションのサーバー

実行方法
- ターミナル1で "go run <ファイル名>" を実行してサーバーを立ち上げる
- Webブラウザで "http://localhost:8080/todo" にアクセス
*/

package main

import (
	"html/template"
	"log"
	"net/http"
)

var todoList []string

func handleTodo(w http.ResponseWriter, r *http.Request){
	// todo.htmlをテンプレートファイルとして読み込み、tに代入
	t, _ := template.ParseFiles("templates/todo.html")
	// 最終的なHTMLファイルを生成して返却
	// w http.ResponseWriter ： HTTPレスポンスとして出力
	// todoList : テンプレートファイルに適用するデータ
	t.Execute(w, todoList)
}

func main(){
	todoList = append(todoList, "顔を洗う", "朝食を食べる", "歯を磨く")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/todo", handleTodo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}