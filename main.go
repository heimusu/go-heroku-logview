// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		log.Fatal("$PORT must be set")
// 	}

// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":"+port, nil)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", r.URL.Path[1:])
// }

package main

import (
	"os"

	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		// log.Fatal("$PORT must be set")
		port = ":1323"
	}

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/hello", handler.MainPage())

	// サーバー起動
	// e.Start(":1323") //ポート番号指定してね
	e.Start(port)
}
