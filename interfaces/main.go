package main

import (
	"fs"
)
import _ "fops/infrastructure/repository"

func main() {
	fs.Run("FOPS")
	/*	http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		})*/
}
