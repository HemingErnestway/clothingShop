package storage

import (
	"clothingShop/entity"
	"sync"
)

type AgeGroupMx struct {
	mtx       sync.RWMutex
	iter      uint32
	ageGroups map[uint32]entity.AgeGroup
}

var ageGroupMx AgeGroupMx

func init() {
	ageGroupMx = AgeGroupMx{
		ageGroups: make(map[uint32]entity.AgeGroup),
	}
}

func AgeGroupCreate(ageGroup entity.AgeGroup) *entity.AgeGroup {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	ageGroupMx.iter++
	ageGroup.Uuid = ageGroupMx.iter
	ageGroupMx.ageGroups[ageGroupMx.iter] = ageGroup

	return &ageGroup
}

func AgeGroupRead(id uint32) *entity.AgeGroup {
	ageGroupMx.mtx.RLock()
	defer ageGroupMx.mtx.RUnlock()

	if el, ok := ageGroupMx.ageGroups[id]; ok {
		return &el
	}

	return nil
}

func AgeGroupsRead() []entity.AgeGroup {
	ageGroupMx.mtx.RLock()
	defer ageGroupMx.mtx.RUnlock()

	ageGroupList := make([]entity.AgeGroup, len(ageGroupMx.ageGroups))
	iter := 0
	for key := range ageGroupMx.ageGroups {
		ageGroupList[iter] = ageGroupMx.ageGroups[key]
		iter++
	}

	return ageGroupList
}

func AgeGroupUpdate(new entity.AgeGroup, id uint32) *entity.AgeGroup {
	ageGroupMx.mtx.Lock()
	defer ageGroupMx.mtx.Unlock()

	current := ageGroupMx.ageGroups[id]

	if new.Name != "" {
		current.Name = new.Name
	}

	ageGroupMx.ageGroups[id] = current
	return &current
}

func AgeGroupDelete(id uint32) string {
	ageGroupMx.mtx.Lock()
	defer ageGroupMx.mtx.Unlock()

	delete(ageGroupMx.ageGroups, id)

	return "successfully deleted"
}
