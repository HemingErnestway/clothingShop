//package storage
//
//import (
//	"clothingShop/entity"
//	"github.com/google/uuid"
//)
//
//var users map[string]entity.User
//
//func init() {
//	users = make(map[string]entity.User, 0)
//}
//
//func UserNew(u entity.User) *entity.User {
//	u.Uuid = uuid.New()
//	users[u.Login] = u
//	return &u
//}
//
//func UserFind(login string) *entity.User {
//	u := users[login]
//	return &u
//}

package storage

import (
	"clothingShop/entity"
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
	defer userMx.mtx.Lock()
	userMx.iter++
	user.Uuid = userMx.iter
	userMx.users[userMx.iter] = user
	return &user
}

func UserGetAll() []entity.User {
	userMx.mtx.RLocker()
	defer userMx.mtx.RUnlock()
	lst := make([]entity.User, len(userMx.users))
	iter := 0
	for key := range userMx.users {
		lst[iter] = userMx.users[key]
		iter++
	}
	return lst
}

func UserGet(uid uint32) *entity.User {
	userMx.mtx.RLock()
	defer userMx.mtx.RUnlock()
	if el, ok := userMx.users[uid]; ok {
		return &el
	}
	return nil
}
