package main

import (
	"bytes"
	"fmt"
)

// 書かれた内容を記憶しておくバッファ
// Write() で書き込まれた内容をためて置いてあとでまとめて結果を
// 受け取る bytes.Buffer
func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example"))
	fmt.Println(buffer.String())
}
