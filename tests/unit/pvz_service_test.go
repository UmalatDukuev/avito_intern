package unit

import (
	"errors"
	"testing"
	"time"

	"avito_intern/internal/handler/entity"
	"avito_intern/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPVZRepository struct {
	mock.Mock
}

func (m *MockPVZRepository) CreatePVZ(pvz entity.PVZ) (string, error) {
	args := m.Called(pvz)
	return args.String(0), args.Error(1)
}

func (m *MockPVZRepository) GetPVZWithDetails(startDate, endDate *time.Time, page, limit int) ([]entity.PVZResponse, error) {
	args := m.Called(startDate, endDate, page, limit)
	if res, ok := args.Get(0).([]entity.PVZResponse); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPVZRepository) GetByID(pvzID string) (*entity.PVZ, error) {
	args := m.Called(pvzID)
	if pvz := args.Get(0); pvz != nil {
		return pvz.(*entity.PVZ), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestPVZService_GetByID(t *testing.T) {
	mockRepo := new(MockPVZRepository)
	pvzSvc := service.NewPVZService(mockRepo)
	expectedPVZ := &entity.PVZ{ID: "123", City: "Москва", RegistrationDate: time.Now()}
	mockRepo.On("GetByID", "123").Return(expectedPVZ, nil)
	pvz, err := pvzSvc.GetByID("123")
	assert.NoError(t, err)
	assert.Equal(t, expectedPVZ, pvz)
	mockRepo.On("GetByID", "notfound").Return(nil, nil)
	pvz, err = pvzSvc.GetByID("notfound")
	assert.NoError(t, err)
	assert.Nil(t, pvz)
	mockRepo.On("GetByID", "error").Return(nil, errors.New("db error"))
	pvz, err = pvzSvc.GetByID("error")
	assert.Error(t, err)
	assert.Nil(t, pvz)
}
