package main

import (
	"gateway/internal/boot"
	"gateway/web"
)

func main() {
	boot.New().Boot()

	web.Serve()
}
