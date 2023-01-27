package main

import (
	"order/internal/boot"
	"order/internal/workers"
)

func main() {
	boot.New().Boot()

	workers.Run()
}
