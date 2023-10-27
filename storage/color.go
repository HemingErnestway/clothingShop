package storage

import (
	"clothingShop/entity"
	"sync"
)

type ColorMx struct {
	mtx    sync.RWMutex
	iter   uint32
	colors map[uint32]entity.Color
}

var colorMx ColorMx

func init() {
	colorMx = ColorMx{
		colors: make(map[uint32]entity.Color),
	}
}

func ColorCreate(color entity.Color) *entity.Color {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	colorMx.iter++
	color.Uuid = colorMx.iter
	colorMx.colors[colorMx.iter] = color

	return &color
}

func ColorRead(id uint32) *entity.Color {
	colorMx.mtx.RLock()
	defer colorMx.mtx.RUnlock()

	if el, ok := colorMx.colors[id]; ok {
		return &el
	}

	return nil
}

func ColorsRead() []entity.Color {
	colorMx.mtx.RLock()
	defer colorMx.mtx.RUnlock()

	colorList := make([]entity.Color, len(colorMx.colors))
	iter := 0
	for key := range colorMx.colors {
		colorList[iter] = colorMx.colors[key]
		iter++
	}

	return colorList
}

func ColorUpdate(new entity.Color, id uint32) *entity.Color {
	colorMx.mtx.Lock()
	defer colorMx.mtx.Unlock()

	current := colorMx.colors[id]

	if new.Name != "" {
		current.Name = new.Name
	}

	colorMx.colors[id] = current
	return &current
}

func ColorDelete(id uint32) string {
	colorMx.mtx.Lock()
	defer colorMx.mtx.Unlock()
	delete(colorMx.colors, id)

	return "successfully deleted"
}
