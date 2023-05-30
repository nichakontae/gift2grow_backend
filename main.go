package main

import (
	"gift2grow_backend/loaders/fiber"
	"gift2grow_backend/loaders/mysql"
)

func main() {
	mysql.Init()
	fiber.Init()
}
