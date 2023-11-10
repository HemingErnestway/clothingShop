package storage

import (
	"clothingShop/db"
	"clothingShop/entity"
)

func SeasonCreate(season entity.Season) *entity.Season {
	db.DB().Create(&season)
	return &season
}

func SeasonRead(id uint32) *entity.Season {
	var season entity.Season
	db.DB().Table(season.TableName()).Where(
		"uuid = ?", id).Find(&season)
	return &season
}

func SeasonsRead() []*entity.Season {
	var seasons []*entity.Season
	db.DB().Find(&seasons)
	return seasons
}

func SeasonUpdate(new entity.Season, id uint32) *entity.Season {
	var current entity.Season
	db.DB().Table(current.TableName()).Where(
		"uuid = ?", id).Find(&current)

	if new.Name != "" {
		current.Name = new.Name
	}

	db.DB().Save(&current)
	return SeasonRead(id)
}

func SeasonDelete(id uint32) string {
	var season entity.Season
	db.DB().Table(season.TableName()).Where(
		"uuid = ?", id).Delete(&season)
	return "successfully deleted"
}
