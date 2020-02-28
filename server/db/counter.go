package db

import (
	"github.com/jinzhu/gorm"
)

type CheckinCounter struct {
	gorm.Model
	Counter int
}

type CounterStore struct {
	db *gorm.DB
}

func NewCheckinCounterStore(db *gorm.DB) *CounterStore {
	return &CounterStore{
		db: db,
	}
}

func (cs *CounterStore) GetCounter() (*CheckinCounter, error) {
	var m CheckinCounter

	cs.db.FirstOrCreate(&m, 1)

	return &m, nil
}

func (cs *CounterStore) UpdateCounter() (*CheckinCounter, error) {
	var m CheckinCounter

	cs.db.FirstOrCreate(&m, 1)
	currentVal := m.Counter
	currentVal++

	cs.db.Model(&m).Update("Counter", currentVal)

	return &m, nil
}