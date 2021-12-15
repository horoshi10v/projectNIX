package main

import (
	"fmt"
	"projectNIX/repositories/filesystem"
)

func main() {

	repo := filesystem.UserFileRepository{}

	user := repo.GetByEmail("")
	fmt.Println(user)

}
