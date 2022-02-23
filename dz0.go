package main

import (
	"github.com/kyokomi/emoji/v2"
)

func main() {
	emoji.Println(hello(":world_map:!"))
}

func hello(name string) string {
	return "Hello, " + name
}
