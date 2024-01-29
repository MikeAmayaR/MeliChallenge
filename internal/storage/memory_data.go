package storage

import (
	"sync"

	"github.com/MikeAmayaR/MeliChallenge.git/internal/model"
)

type InMemoryStorage struct {
	mu             sync.Mutex
	satellitesData map[string]model.Satellite
}

var memoryStorage = InMemoryStorage{
	satellitesData: make(map[string]model.Satellite),
}

func (s *InMemoryStorage) SaveSatelliteData(satellite model.Satellite) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.satellitesData[satellite.Name] = satellite
}

func (s *InMemoryStorage) RetrieveAllSatelliteData() []model.Satellite {
	s.mu.Lock()
	defer s.mu.Unlock()
	dataSlice := make([]model.Satellite, 0, len(s.satellitesData))
	for _, data := range s.satellitesData {
		dataSlice = append(dataSlice, data)
	}
	return dataSlice
}

func (s *InMemoryStorage) HasSufficientData() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.satellitesData) >= 3
}

func (s *InMemoryStorage) ClearData() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.satellitesData = make(map[string]model.Satellite)
}

func GetMemoryStorage() *InMemoryStorage {
	return &memoryStorage
}
