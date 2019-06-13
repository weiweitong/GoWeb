package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName02(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	// 解析url传递的参数，对于POST则是解析响应包的主体(request body)

	// 这些信息是输出到服务端的打印信息
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello astaxie!")
	// 这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method)
	// 获取请求的方法

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./HTTP/data/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main()  {
	// 设置访问的路由
	http.HandleFunc("/", sayHelloName02)

	// 设置访问的路由
	http.HandleFunc("/login", login)

	// 设置监听的端口
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
