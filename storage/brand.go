package storage

import (
	"clothingShop/entity"
	"sync"
)

type BrandMx struct {
	mtx    sync.RWMutex
	iter   uint32
	brands map[uint32]entity.Brand
}

var brandMx BrandMx

func init() {
	brandMx = BrandMx{
		brands: make(map[uint32]entity.Brand),
	}
}

func BrandCreate(brand entity.Brand) *entity.Brand {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	brandMx.iter++
	brand.Uuid = brandMx.iter
	brandMx.brands[brandMx.iter] = brand

	return &brand
}

func BrandRead(id uint32) *entity.Brand {
	brandMx.mtx.RLock()
	defer brandMx.mtx.RUnlock()

	if el, ok := brandMx.brands[id]; ok {
		return &el
	}

	return nil
}

func BrandsRead() []entity.Brand {
	brandMx.mtx.RLock()
	defer brandMx.mtx.RUnlock()

	brandList := make([]entity.Brand, len(brandMx.brands))
	iter := 0
	for key := range brandMx.brands {
		brandList[iter] = brandMx.brands[key]
		iter++
	}

	return brandList
}

func BrandUpdate(new entity.Brand, id uint32) *entity.Brand {
	brandMx.mtx.Lock()
	defer brandMx.mtx.Unlock()

	current := brandMx.brands[id]

	if new.Name != "" {
		current.Name = new.Name
	}

	brandMx.brands[id] = current
	return &current
}

func BrandDelete(id uint32) string {
	brandMx.mtx.Lock()
	defer brandMx.mtx.Unlock()
	delete(brandMx.brands, id)

	return "successfully deleted"
}
