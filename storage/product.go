package storage

import (
	"clothingShop/db"
	"clothingShop/dto"
	"clothingShop/entity"
	"fmt"
	"reflect"
)

func ProductCreate(product entity.Product) *entity.Product {
	db.DB().Create(&product)
	return &product
}

func ProductRead(id uint32) *entity.Product {
	var product entity.Product
	db.DB().Table(product.TableName()).Where(
		"uuid = ?", id).Find(&product)
	return &product
}

func ProductsRead() []*entity.Product {
	var products []*entity.Product
	db.DB().Find(&products)
	return products
}

func ProductUpdate(new dto.Product, id uint32) *entity.Product {
	var current entity.Product
	db.DB().Table(current.TableName()).Where(
		"uuid = ?", id).Find(&current)

	currentStruct := reflect.ValueOf(&current).Elem()
	newStruct := reflect.ValueOf(&new.Product).Elem()

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

	db.DB().Save(&current)
	return ProductRead(id)
}

func ProductDelete(id uint32) string {
	var product entity.Product
	db.DB().Table(product.TableName()).Where(
		"uuid = ?", id).Delete(&product)
	return "successfully deleted"
}
