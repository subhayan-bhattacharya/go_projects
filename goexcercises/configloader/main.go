package main

import (
	"configloader/config"
	"fmt"
	"os"
)

func main() {
	os.Setenv("PORT", "8080")
	os.Setenv("DATABASE_URL", "postgres://localhost:5432/db")
	os.Setenv("TIMEOUT", "30s")
	os.Setenv("DEBUG", "true")
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("The loaded config %+v\n", *cfg)
}
