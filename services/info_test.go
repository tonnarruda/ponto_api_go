package services_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tonnarruda/ponto_api_go/services"
	"github.com/tonnarruda/ponto_api_go/structs"
)

// MockInfoRepository is a mock implementation of the InfoRepository interface
type MockInfoRepository struct {
	mock.Mock
}

func (m *MockInfoRepository) GetAll() ([]structs.Info, error) {
	args := m.Called()
	return args.Get(0).([]structs.Info), args.Error(1)
}

func TestGetAllInfo_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockInfoRepository)
	infoService := services.NewInfoService(mockRepo)
	expectedInfos := []structs.Info{
		{VersaoBD: 307, StatusBD: "OK", Sistema: "Ponto", VersaoBDBeta: 0, Atualizando: 0, Fortes: 0, ConvertePonto3: 0},
	}

	mockRepo.On("GetAll").Return(expectedInfos, nil)

	// Act
	infos, err := infoService.GetAllInfo()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedInfos, infos)
	mockRepo.AssertExpectations(t)
}

func TestGetAllInfo_Error(t *testing.T) {
	// Arrange
	mockRepo := new(MockInfoRepository)
	infoService := services.NewInfoService(mockRepo)
	expectedError := errors.New("some error")

	// Retornando um slice vazio em vez de nil
	mockRepo.On("GetAll").Return([]structs.Info{}, expectedError)

	// Act
	infos, err := infoService.GetAllInfo()

	// Assert
	assert.Error(t, err)
	assert.NotNil(t, infos)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
