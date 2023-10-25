package storage

import (
	"clothingShop/entity"
	"sync"
)

type CategoryMx struct {
	mtx        sync.RWMutex
	iter       uint32
	categories map[uint32]entity.Category
}

var categoryMx CategoryMx

func init() {
	categoryMx = CategoryMx{
		categories: make(map[uint32]entity.Category),
	}
}

func CategoryCreate(category entity.Category) *entity.Category {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	categoryMx.iter++
	category.Uuid = categoryMx.iter
	categoryMx.categories[categoryMx.iter] = category

	return &category
}

func CategoryRead(id uint32) *entity.Category {
	categoryMx.mtx.RLock()
	defer categoryMx.mtx.RUnlock()

	if el, ok := categoryMx.categories[id]; ok {
		return &el
	}

	return nil
}

func CategoriesRead() []entity.Category {
	categoryMx.mtx.RLock()
	defer categoryMx.mtx.RUnlock()

	categoryList := make([]entity.Category, len(categoryMx.categories))
	iter := 0
	for key := range categoryMx.categories {
		categoryList[iter] = categoryMx.categories[key]
		iter++
	}

	return categoryList
}

func CategoryUpdate(new entity.Category, id uint32) *entity.Category {
	categoryMx.mtx.Lock()
	defer categoryMx.mtx.Unlock()

	current := categoryMx.categories[id]

	if new.Name != "" {
		current.Name = new.Name
	}

	categoryMx.categories[id] = current
	return &current
}

func CategoryDelete(id uint32) string {
	categoryMx.mtx.Lock()
	defer categoryMx.mtx.Unlock()
	delete(categoryMx.categories, id)

	return "successfully deleted"
}
