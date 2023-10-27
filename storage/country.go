package storage

import (
	"clothingShop/entity"
	"sync"
)

type CountryMx struct {
	mtx       sync.RWMutex
	iter      uint32
	countries map[uint32]entity.Country
}

var countryMx CountryMx

func init() {
	countryMx = CountryMx{
		countries: make(map[uint32]entity.Country),
	}
}

func CountryCreate(country entity.Country) *entity.Country {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	countryMx.iter++
	country.Uuid = countryMx.iter
	countryMx.countries[countryMx.iter] = country

	return &country
}

func CountryRead(id uint32) *entity.Country {
	countryMx.mtx.RLock()
	defer countryMx.mtx.RUnlock()

	if el, ok := countryMx.countries[id]; ok {
		return &el
	}

	return nil
}

func CountriesRead() []entity.Country {
	countryMx.mtx.RLock()
	defer countryMx.mtx.RUnlock()

	countryList := make([]entity.Country, len(countryMx.countries))
	iter := 0
	for key := range countryMx.countries {
		countryList[iter] = countryMx.countries[key]
		iter++
	}

	return countryList
}

func CountryUpdate(new entity.Country, id uint32) *entity.Country {
	countryMx.mtx.Lock()
	defer countryMx.mtx.Unlock()

	current := countryMx.countries[id]

	if new.Name != "" {
		current.Name = new.Name
	}

	countryMx.countries[id] = current
	return &current
}

func CountryDelete(id uint32) string {
	countryMx.mtx.Lock()
	defer countryMx.mtx.Unlock()
	delete(countryMx.countries, id)

	return "successfully deleted"
}
