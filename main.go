package zadan0

import "github.com/kyokomi/emoji/v2"

func Main() {
	emoji.Println(Hello(":world_map:!"))
}

func Hello(name string) string {
	return "Hello, " + name
}
