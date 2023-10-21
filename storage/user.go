package storage

import (
	"clothingShop/dto"
	"clothingShop/entity"
	"fmt"
	"reflect"
	"sync"
)

type UserMx struct {
	mtx   sync.RWMutex
	iter  uint32
	users map[uint32]entity.User
}

var userMx UserMx

func init() {
	userMx = UserMx{
		users: make(map[uint32]entity.User),
	}
}

func UserCreate(user entity.User) *entity.User {
	userMx.mtx.Lock()
	defer userMx.mtx.Unlock()

	userMx.iter++
	user.Uuid = userMx.iter
	userMx.users[userMx.iter] = user

	return &user
}

func UserRead(id uint32) *entity.User {
	userMx.mtx.RLock()
	defer userMx.mtx.RUnlock()

	if el, ok := userMx.users[id]; ok {
		return &el
	}

	return nil
}

func UsersRead() []entity.User {
	userMx.mtx.RLock()
	defer userMx.mtx.RUnlock()

	userList := make([]entity.User, len(userMx.users))
	iter := 0
	for key := range userMx.users {
		userList[iter] = userMx.users[key]
		iter++
	}

	return userList
}

func UserUpdate(new dto.User, id uint32) *entity.User {
	userMx.mtx.Lock()
	defer userMx.mtx.Unlock()

	current := userMx.users[id]

	currentStruct := reflect.ValueOf(&current).Elem()
	newStruct := reflect.ValueOf(&new.User).Elem()

	for i := 0; i < newStruct.NumField(); i++ {
		newField := newStruct.Type().Field(i)
		if !newStruct.FieldByName(newField.Name).IsZero() {
			fmt.Println(newField.Name)
			currentField := currentStruct.FieldByName(newField.Name)
			currentField.Set(newStruct.FieldByName(newField.Name))
		}
	}

	for _, attr := range new.ClearAttr {
		s := reflect.ValueOf(&current).Elem()
		field := s.FieldByName(attr)
		if field.CanSet() {
			field.SetZero()
		}
	}

	userMx.users[id] = current
	return &current
}

func UserDelete(id uint32) string {
	userMx.mtx.Lock()
	defer userMx.mtx.Unlock()

	delete(userMx.users, id)

	return "successfully deleted"
}
