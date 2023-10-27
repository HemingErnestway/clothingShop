package storage

import (
	"clothingShop/entity"
	"sync"
)

type SeasonMx struct {
	mtx     sync.RWMutex
	iter    uint32
	seasons map[uint32]entity.Season
}

var seasonMx SeasonMx

func init() {
	seasonMx = SeasonMx{
		seasons: make(map[uint32]entity.Season),
	}
}

func SeasonCreate(season entity.Season) *entity.Season {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	seasonMx.iter++
	season.Uuid = seasonMx.iter
	seasonMx.seasons[seasonMx.iter] = season

	return &season
}

func SeasonRead(id uint32) *entity.Season {
	seasonMx.mtx.RLock()
	defer seasonMx.mtx.RUnlock()

	if el, ok := seasonMx.seasons[id]; ok {
		return &el
	}

	return nil
}

func SeasonsRead() []entity.Season {
	seasonMx.mtx.RLock()
	defer seasonMx.mtx.RUnlock()

	seasonList := make([]entity.Season, len(seasonMx.seasons))
	iter := 0
	for key := range seasonMx.seasons {
		seasonList[iter] = seasonMx.seasons[key]
		iter++
	}

	return seasonList
}

func SeasonUpdate(new entity.Season, id uint32) *entity.Season {
	seasonMx.mtx.Lock()
	defer seasonMx.mtx.Unlock()

	current := seasonMx.seasons[id]

	if new.Name != "" {
		current.Name = new.Name
	}

	seasonMx.seasons[id] = current
	return &current
}

func SeasonDelete(id uint32) string {
	seasonMx.mtx.Lock()
	defer seasonMx.mtx.Unlock()
	delete(seasonMx.seasons, id)

	return "successfully deleted"
}
