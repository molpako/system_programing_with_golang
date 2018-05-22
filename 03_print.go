package main

import (
	"os"
)

func main() {
	// fmt.Println では最終的に os.Stdout.Write を呼び出している
	os.Stdout.Write([]byte("os.Stdout example"))
}
