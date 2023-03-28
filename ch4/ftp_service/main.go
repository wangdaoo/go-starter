package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// 在终端中使用命令 go build main.go 来编译生成可执行文件。
// 使用命令 ./main 来运行程序，此时服务端会监听本地8000端口，等待客户端连接。
// 在另一个终端中使用 telnet localhost 8000 命令来连接服务端，然后就可以使用以下几个命令进行文件传输：
// ls: 列出当前目录下的所有文件和文件夹。
// cd [directory]: 切换到指定的目录。
// get [filename]: 下载指定的文件。
// send [filename]: 上传指定的文件。
// close: 关闭连接。

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		cmd := input.Text()
		if cmd == "close" {
			fmt.Fprintln(conn, "Goodbye!")
			return
		}
		args := strings.Split(cmd, " ")
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(conn, "usage: cd directory")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Fprintln(conn, "error:", err)
				continue
			}
			fmt.Fprintln(conn, "OK")
		case "ls":
			files, err := os.ReadDir(".")
			if err != nil {
				fmt.Fprintln(conn, "error:", err)
				continue
			}
			for _, file := range files {
				fmt.Fprintln(conn, file.Name())
			}
		case "get":
			if len(args) < 2 {
				fmt.Fprintln(conn, "usage: get filename")
				continue
			}
			file, err := os.Open(args[1])
			if err != nil {
				fmt.Fprintln(conn, "error:", err)
				continue
			}
			_, err = io.Copy(conn, file)
			if err != nil {
				fmt.Fprintln(conn, "error:", err)
				continue
			}
			file.Close()
		case "send":
			if len(args) < 2 {
				fmt.Fprintln(conn, "usage: send filename")
				continue
			}
			file, err := os.Create(args[1])
			if err != nil {
				fmt.Fprintln(conn, "error:", err)
				continue
			}
			_, err = io.Copy(file, conn)
			if err != nil {
				fmt.Fprintln(conn, "error:", err)
				continue
			}
			file.Close()
		default:
			fmt.Fprintln(conn, "unknown command:", args[0])
		}
	}
}
