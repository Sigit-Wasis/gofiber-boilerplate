package seed

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Sigit-Wasis/gofiber-boilerplate/internal/db"
)


func Run() error {
d := db.Get()
if d == nil {
return fmt.Errorf("db not initialized")
}


seedFile := filepath.Join("migrations", "seed_users.sql")
b, err := os.ReadFile(seedFile)
if err != nil {
return err
}


if _, err = d.Exec(string(b)); err != nil {
return err
}


return nil
}