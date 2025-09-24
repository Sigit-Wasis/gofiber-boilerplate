package main

import (
	"log"

	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/config"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/db"
	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/seed"
)


func main() {
cfg := config.Load()
if err := db.Init(cfg.DatabaseURL); err != nil {
log.Fatalf("db init: %v", err)
}
defer db.Close()


if err := seed.Run(); err != nil {
log.Fatalf("seed error: %v", err)
}
log.Println("seed completed")
}