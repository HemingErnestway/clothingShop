package storage

import (
	"clothingShop/dto"
	"clothingShop/entity"
	"reflect"
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
	product.DiscountPercent = calculateDiscount(&product)
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

func ProductUpdate(new dto.Product, id uint32) *entity.Product {
	productMx.mtx.Lock()
	defer productMx.mtx.Unlock()

	current := productMx.products[id]

	currentStruct := reflect.ValueOf(&current).Elem()
	newStruct := reflect.ValueOf(&new.Product).Elem()

	for i := 0; i < newStruct.NumField(); i++ {
		newField := newStruct.Type().Field(i)
		if !newStruct.FieldByName(newField.Name).IsZero() {
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

	current.DiscountPercent = calculateDiscount(&current)

	productMx.products[id] = current
	return &current
}

func ProductDelete(id uint32) string {
	productMx.mtx.Lock()
	defer productMx.mtx.Unlock()

	delete(productMx.products, id)

	return "successfully deleted"
}

func calculateDiscount(product *entity.Product) uint8 {
	if product.NewPrice == 0.0 {
		return 0
	}
	return uint8(100 * (product.Price - product.NewPrice) / product.Price)
}
