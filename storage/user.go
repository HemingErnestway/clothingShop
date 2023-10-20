package storage

import (
	"clothingShop/dto"
	"clothingShop/entity"
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
	// TODO: consider refactoring using reflect
	switch {
	case new.Name != "":
		current.Name = new.Name
	case new.Surname != "":
		current.Surname = new.Surname
	case new.Email != "":
		current.Email = new.Email
	case new.Address != "":
		current.Address = new.Address
	case new.BonusPoints != 0:
		current.BonusPoints = new.BonusPoints
	case new.BirthDate != "":
		current.BirthDate = new.BirthDate
	case new.Login != "":
		current.Login = new.Login
	case new.Password != "":
		current.Password = new.Password
	case new.Access != 0:
		current.Access = new.Access
	}

	for _, attr := range new.ClearAttr {
		s := reflect.ValueOf(&current).Elem()
		if s.Kind() == reflect.Struct {
			field := s.FieldByName(attr)
			if field.CanSet() {
				field.SetZero()
			}
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
