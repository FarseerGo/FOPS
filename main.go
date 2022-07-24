package main

import (
	"fmt"
	_ "fops/infrastructure/repository"
	"fs"
	"fs/configure"
)

func main() {
	fs.Run("FOPS")
	fmt.Println(configure.GetString("Database.fops"))
}
