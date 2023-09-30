package storage

import (
	"clothingShop/entity"
	"github.com/google/uuid"
)

var users map[string]entity.User

func init() {
	users = make(map[string]entity.User, 0)
}

func UserNew(u entity.User) *entity.User {
	u.Uuid = uuid.New()
	users[u.Login] = u
	return &u
}

func UserFind(login string) *entity.User {
	u := users[login]
	return &u
}
