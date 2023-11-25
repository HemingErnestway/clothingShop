package storage

import (
	"clothingShop/db"
	"clothingShop/entity"
	"fmt"
	"reflect"
)

func TokenCreate(token entity.Token) *entity.Token {
	db.DB().Create(&token)
	return &token
}

func TokenRead(id uint32) *entity.Token {
	var token entity.Token
	db.DB().Table(token.TableName()).Where(
		"uuid = ?", id).Find(&token)
	return &token
}

func TokensRead() []*entity.Token {
	var tokens []*entity.Token
	db.DB().Find(&tokens)
	return tokens
}

func TokenUpdate(new entity.Token, id uint32) *entity.Token {
	var current entity.Token
	db.DB().Table(current.TableName()).Where(
		"uuid = ?", id).Find(&current)

	currentStruct := reflect.ValueOf(&current).Elem()
	newStruct := reflect.ValueOf(&new.Token).Elem()

	for i := 0; i < newStruct.NumField(); i++ {
		newField := newStruct.Type().Field(i)
		if !newStruct.FieldByName(newField.Name).IsZero() {
			fmt.Println(newField.Name)
			currentField := currentStruct.FieldByName(newField.Name)
			currentField.Set(newStruct.FieldByName(newField.Name))
		}
	}

	db.DB().Save(&current)
	return TokenRead(id)
}

func TokenDelete(id uint32) string {
	var token entity.Token
	db.DB().Table(token.TableName()).Where(
		"uuid = ?", id).Delete(&token)
	return "successfully deleted"
}
