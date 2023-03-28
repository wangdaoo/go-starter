package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 输出请求的 URL 路径
// func handler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "URL.Path111 = %q\n", r.URL.Path)
// 	// 输出一个 h1 标签
// 	fmt.Fprintf(w, "<h1>URL.Path = %q</h1>", r.URL.Path)
// }

// 输出 HTTP request 的 header
func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	// 打印本机 IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, address := range addrs {
		// 检查 ip 地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// fmt.Println(ipnet.IP.String())
				fmt.Fprintf(w, "Local IP = %q\n", ipnet.IP.String())
			}
		}
	}

	// 获取 url 中的参数
	// http://localhost:8000/?name=liang&age=18
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	fmt.Fprintf(w, "r.URL.RawQuery = %q\n", r.URL.RawQuery)

	// 获取 url 中的参数
	// http://localhost:8000/?name=liang&age=18
	fmt.Fprintf(w, "r.URL.Query() = %q\n", r.URL.Query())

	// 获取 url 中的参数
	// http://localhost:8000/?name=liang&age=18
	fmt.Fprintf(w, "r.URL.Query().Get(\"name\") = %q\n", r.URL.Query().Get("name"))

	// 获取 url 中的参数
	// http://localhost:8000/?name=liang&age=18
	fmt.Fprintf(w, "r.URL.Query().Get(\"age\") = %q\n", r.URL.Query().Get("age"))

	// 获取 url 中的参数, 产生一个 map
	// http://localhost:8000/?name=liang&age=18
	fmt.Fprintf(w, "r.URL.Query()[\"name\"] = %q\n", r.URL.Query()["name"])

	x := 1
	p := &x
	fmt.Println(*p) // "1"
	fmt.Println(p)  // 变量 p 的内存地址
}
