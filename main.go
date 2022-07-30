package main

import (
	"fmt"
	_ "fops/infrastructure/repository"
	"github.com/farseernet/farseer.go/configure"
	"github.com/farseernet/farseer.go/fsApp"
)

func main() {
	fsApp.Initialize("FOPS")
	fmt.Println(configure.GetString("Database.fops"))
}
