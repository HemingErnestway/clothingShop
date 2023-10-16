package storage

import (
	"clothingShop/entity"
	"sync"
)

type ProductMx struct {
	mtx      sync.RWMutex
	iter     uint32
	products map[uint32]entity.Product
}

var productMx ProductMx

func init() {
	productMx = ProductMx{
		products: make(map[uint32]entity.Product),
	}
}

func ProductCreate(product entity.Product) *entity.Product {
	productMx.mtx.Lock()
	defer productMx.mtx.Unlock()

	productMx.iter++
	product.Uuid = productMx.iter
	productMx.products[productMx.iter] = product

	return &product
}

func ProductRead(id uint32) *entity.Product {
	productMx.mtx.RLock()
	defer productMx.mtx.RUnlock()

	if el, ok := productMx.products[id]; ok {
		return &el
	}

	return nil
}

func ProductsRead() []entity.Product {
	productMx.mtx.RLock()
	defer productMx.mtx.RUnlock()

	productList := make([]entity.Product, len(productMx.products))
	iter := 0
	for key := range productMx.products {
		productList[iter] = productMx.products[key]
		iter++
	}

	return productList
}

func ProductUpdate(new entity.Product, id uint32) *entity.Product {
	productMx.mtx.Lock()
	defer productMx.mtx.Unlock()

	current := productMx.products[id]
	// TODO: consider refactoring using reflect
	switch {
	case new.CategoryId != 0:
		current.CategoryId = new.CategoryId
	case new.ColorId != 0:
		current.ColorId = new.ColorId
	case new.SeasonId != 0:
		current.SeasonId = new.SeasonId
	case new.SizeId != 0:
		current.SizeId = new.SizeId
	case new.ManufacturerId != 0:
		current.ManufacturerId = new.ManufacturerId
	case new.BrandId != 0:
		current.BrandId = new.BrandId
	case new.GenderId != 0:
		current.GenderId = new.GenderId
	case new.AgeGroupId != 0:
		current.AgeGroupId = new.AgeGroupId
	case new.PriceRoubles != 0.0:
		current.PriceRoubles = new.PriceRoubles
	}

	productMx.products[id] = current
	return &current
}

func ProductDelete(id uint32) []entity.Product {
	productMx.mtx.Lock()
	defer productMx.mtx.Unlock()

	delete(productMx.products, id)

	productList := make([]entity.Product, len(productMx.products))
	iter := 0
	for key := range productMx.products {
		productList[iter] = productMx.products[key]
		iter++
	}

	return productList
}
