package main

import (
	"encoding/json"
	"os"
)

func main() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	data := map[string]string{
		"example": "encodeing/json",
		"hello":   "world",
	}
	encoder.Encode(data)
}
