package services

import (
	"github.com/tonnarruda/ponto_api_go/repositories"
	"github.com/tonnarruda/ponto_api_go/structs"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) CreateUser(user *structs.Usuario) error {
	err := s.userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetAllUsers() ([]structs.UserResponse, error) {
	return s.userRepository.GetAll()
}

func (s *UserService) GetUserByCodigo(codigo string) (*structs.UserResponse, error) {
	return s.userRepository.GetByCodigo(codigo)
}
