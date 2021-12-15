package filesystem

import (
	"encoding/json"
	"io"
	"os"
	"projectNIX/repositories/models"
)

type UserFileRepository struct {
}

func (ufu UserFileRepository) GetByEmail(_ string) (user models.User) {

	userData := []byte{}

	file, err := os.Open("./datastore/files/user_1.json")
	if err != nil {
		panic(err)
		return models.User{}
	}
	defer file.Close()
	for {
		chunk := make([]byte, 1024)
		_, err = file.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		userData = append(userData, chunk...)
	}

	err = json.Unmarshal(userData, &user)
	if err != nil {
		panic(err)
	}

	return user
}
