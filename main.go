package main

import (
	"gift2grow_backend/loaders/fiber"
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/storage"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	mysql.Init()
	fiber.Init()
	storage.Init()
}
