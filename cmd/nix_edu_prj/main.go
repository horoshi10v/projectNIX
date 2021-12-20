package main

import (
	"fmt"
	"projectNIX/pkg/repositories/filesystem"
)

func main() {

	repo := filesystem.UserFileRepository{}

	user := repo.GetByEmail("")
	fmt.Println(user)

}
