package main

import (
	"io"
	"net"
	"os"
)

func main() {
	// net.Conn という通信のコネクションを表すインタフェースを返す
	// net.Conn は io.Writer, io.Reader のハイブリッドなインタフェース
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0¥r¥nHost: ascii.jp¥r¥n¥r¥n"))
	io.Copy(os.Stdout, conn)

}
