package storage

import (
	"clothingShop/db"
	"clothingShop/dto"
	"clothingShop/entity"
	"fmt"
	"reflect"
)

func UserCreate(user entity.User) *entity.User {
	db.DB().Create(&user)
	return &user
}

func UserRead(id uint32) *entity.User {
	var user entity.User
	db.DB().Table(user.TableName()).Where(
		"uuid = ?", id).Find(&user)
	return &user
}

func UsersRead() []*entity.User {
	var users []*entity.User
	db.DB().Find(&users)
	return users
}

func UserUpdate(new dto.User, id uint32) *entity.User {
	var current entity.User
	db.DB().Table(current.TableName()).Where(
		"uuid = ?", id).Find(&current)

	currentStruct := reflect.ValueOf(&current).Elem()
	newStruct := reflect.ValueOf(&new.User).Elem()

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
	return UserRead(id)
}

func UserDelete(id uint32) string {
	var user entity.User
	db.DB().Table(user.TableName()).Where(
		"uuid = ?", id).Delete(&user)
	return "successfully deleted"
}
