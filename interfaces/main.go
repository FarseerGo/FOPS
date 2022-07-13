package main

import "fs"
import _ "fops/infrastructure/repository"

func main() {
	fs.Run("FOPS")
}
