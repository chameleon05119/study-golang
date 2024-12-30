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

func handleAdd(w http.ResponseWriter, r *http.Request){
	// HTTPリクエストを解析して"todo"という名前のPOSTパラメータから値を取り出し、todoListに追加
	r.ParseForm()
	todo := r.Form.Get("todo")
	todoList = append(todoList, todo)
	// 更新されたtodoListでテンプレートを更新
	handleTodo(w,r)
	// /todoへリダイレクトさせる
	http.Redirect(w, r, "/todo", 303)
}

func main(){
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/todo", handleTodo)

	http.HandleFunc("/add", handleAdd)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}