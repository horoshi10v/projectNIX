package main

import "fmt"
import "github.com/google/uuid"

func main() {
	fmt.Printf("Hello NIX!")
	id := uuid.New()
	fmt.Printf(id.String())
}
