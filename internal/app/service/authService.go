package service

import (
	"fmt"
	"vothings/internal/app/dto"
	"vothings/internal/app/model"
	"vothings/internal/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	Insert(user model.Users) error
	GetByUsername(username string) (model.Users, error)
}

type userService struct {
	userRepository UserRepo
	jwtToken       utils.Jwt
}

// Login implements api.UserService.
func (u *userService) Login(req dto.UserLoginReq) (string, error) {
	user, err := u.userRepository.GetByUsername(req.Username)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", err
	}
	token, err := u.jwtToken.GenerateToken(req.Username, user.Role)

	if err != nil {
		return "", err
	}
	u.jwtToken.VerivyToken(token)

	return token, nil
}

// Register implements api.UserService.
func (u *userService) Register(req dto.UserRegisterReq) error {
	exist, _ := u.userRepository.GetByUsername(req.UserName)

	if exist != (model.Users{}) {
		return fmt.Errorf("the username already exists, please choose another username")
	}

	// if err != nil {
	// 	return err
	// }

	hashedPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)

	if err != nil {
		return err
	}

	user := model.Users{
		FullName: req.FullName,
		Email:    req.Email,
		UserName: req.UserName,
		Password: string(hashedPw),
		Role:     req.Role,
	}
	if err := u.userRepository.Insert(user); err != nil {
		return err
	}
	return nil
}

func NewUserService(repo UserRepo, jwtToken utils.Jwt) *userService {
	return &userService{
		userRepository: repo,
		jwtToken:       jwtToken,
	}
}
