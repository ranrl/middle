package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/ranrl/middle/app"
)

func main() {
	// 中间件
	// http.Handle("/", timeMiddle(http.HandlerFunc(hello)))
	// err := http.ListenAndServe(":8080", nil)
	// fmt.Println(err)

	// httprouter 基础上实现中间件
	router := app.NewRouter()
	router.Use(timeMiddle)
	router.Use(LoginMiddle)
	router.Add("GET", "/", hello)
	router.Add("GET", "/user/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Println("print fe")
		w.Write([]byte("user name:" + p.ByName("name")))
	})
	err := router.Run(":8080")
	fmt.Println(err)

	// httprouter
	// router := httprouter.New()
	// router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 	w.Write([]byte("index"))
	// })
	// // 精确匹配
	// router.GET("/user/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	w.Write([]byte("user name:" + p.ByName("name")))
	// })
	// // 匹配所有, /username/hack/xxx 也能匹配
	// router.GET("/username/*name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	w.Write([]byte("user name:" + p.ByName("name")))
	// })
	// // 设置静态文件目录
	// router.ServeFiles("/static/*filepath", http.Dir("./"))
	// // 捕获异常
	// router.PanicHandler = func(w http.ResponseWriter, r *http.Request, v interface{}) {
	// 	w.Write([]byte("error"))
	// }
	// // not found
	// // router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// // 	w.Write([]byte("<h1>404</h1>"))
	// // })
	// // 自动匹配
	// router.HandleOPTIONS = true
	// router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// 	w.Write([]byte("auto match handler"))
	// })
	// log.Fatal(http.ListenAndServe(":8080", router))

}
func timeMiddle(next httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		timeStart := time.Now()
		log.Println("timeStart")
		next(w, r, p)
		timeElapsed := time.Since(timeStart)
		log.Println(timeElapsed)
	})
}
func LoginMiddle(next httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Println("login success")
		next(w, r, p)
	})
}
func hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("print fe")
	w.Write([]byte("hello"))
}
