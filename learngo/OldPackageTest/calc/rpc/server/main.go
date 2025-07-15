package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	//http://127.0.0.1:8080/add?a=1&b=2
	/**
	1.callID的问题 r.URL.PATH
	2.数据的传输协议 url参数传递的协议
	3.网络的传输协议 http
	4.数据的编码协议 json
	*/
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		fmt.Println("path: ", r.URL.Path)
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{"data": a + b})
		_, _ = w.Write(jData)
	})
	_ = http.ListenAndServe(":8080", nil)
}
