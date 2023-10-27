package storage

import (
	"clothingShop/entity"
	"sync"
)

type SizeMx struct {
	mtx   sync.RWMutex
	iter  uint32
	sizes map[uint32]entity.Size
}

var sizeMx SizeMx

func init() {
	sizeMx = SizeMx{
		sizes: make(map[uint32]entity.Size),
	}
}

func SizeCreate(size entity.Size) *entity.Size {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	sizeMx.iter++
	size.Uuid = sizeMx.iter
	sizeMx.sizes[sizeMx.iter] = size

	return &size
}

func SizeRead(id uint32) *entity.Size {
	sizeMx.mtx.RLock()
	defer sizeMx.mtx.RUnlock()

	if el, ok := sizeMx.sizes[id]; ok {
		return &el
	}

	return nil
}

func SizesRead() []entity.Size {
	sizeMx.mtx.RLock()
	defer sizeMx.mtx.RUnlock()

	sizeList := make([]entity.Size, len(sizeMx.sizes))
	iter := 0
	for key := range sizeMx.sizes {
		sizeList[iter] = sizeMx.sizes[key]
		iter++
	}

	return sizeList
}

func SizeUpdate(new entity.Size, id uint32) *entity.Size {
	sizeMx.mtx.Lock()
	defer sizeMx.mtx.Unlock()

	current := sizeMx.sizes[id]

	if new.Name != "" {
		current.Name = new.Name
	}

	sizeMx.sizes[id] = current
	return &current
}

func SizeDelete(id uint32) string {
	sizeMx.mtx.Lock()
	defer sizeMx.mtx.Unlock()
	delete(sizeMx.sizes, id)

	return "successfully deleted"
}
