package main

import (
	"fmt"
	_ "fops/infrastructure/repository"
	"github.com/farseernet/farseer.go/configure"
	"github.com/farseernet/farseer.go/init"
)

func main() {
	init.Run("FOPS")
	fmt.Println(configure.GetString("Database.fops"))
}
