package filesystem

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"projectNIX/pkg/repositories/models"
)

type UserFileRepository struct {
}

func (ufu UserFileRepository) GetByEmail(_ string) (user models.User) {

	var userData []byte

	file, err := os.Open("./datastore/files/user_1.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for {
		chunk := make([]byte, 64)
		n, err := file.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		userData = append(userData, chunk[:n]...)
	}

	err = json.Unmarshal(userData, &user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}
