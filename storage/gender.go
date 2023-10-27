package storage

import (
	"clothingShop/entity"
	"sync"
)

type GenderMx struct {
	mtx     sync.RWMutex
	iter    uint32
	genders map[uint32]entity.Gender
}

var genderMx GenderMx

func init() {
	genderMx = GenderMx{
		genders: make(map[uint32]entity.Gender),
	}
}

func GenderCreate(gender entity.Gender) *entity.Gender {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	genderMx.iter++
	gender.Uuid = genderMx.iter
	genderMx.genders[genderMx.iter] = gender

	return &gender
}

func GenderRead(id uint32) *entity.Gender {
	genderMx.mtx.RLock()
	defer genderMx.mtx.RUnlock()

	if el, ok := genderMx.genders[id]; ok {
		return &el
	}

	return nil
}

func GendersRead() []entity.Gender {
	genderMx.mtx.RLock()
	defer genderMx.mtx.RUnlock()

	genderList := make([]entity.Gender, len(genderMx.genders))
	iter := 0
	for key := range genderMx.genders {
		genderList[iter] = genderMx.genders[key]
		iter++
	}

	return genderList
}

func GenderUpdate(new entity.Gender, id uint32) *entity.Gender {
	genderMx.mtx.Lock()
	defer genderMx.mtx.Unlock()

	current := genderMx.genders[id]

	if new.Name != "" {
		current.Name = new.Name
	}

	genderMx.genders[id] = current
	return &current
}

func GenderDelete(id uint32) string {
	genderMx.mtx.Lock()
	defer genderMx.mtx.Unlock()
	delete(genderMx.genders, id)

	return "successfully deleted"
}
