package main

import (
	"bufio"
	"os"
)

func main() {
	// 出力結果を一時的にためておいて、ある程度の分量ごとにまとめて書き出す bufio.Writer
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer")

	// Flush を呼ぶと、後続の io.Writer に書き出す
	buffer.Flush()
	buffer.WriteString("example¥n")

	// Flush を呼ばないと書き込まれたデータを持ったまま消滅してしまう
	buffer.Flush()
}
