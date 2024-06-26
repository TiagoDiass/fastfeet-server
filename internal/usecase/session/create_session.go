package usecase

import (
	"errors"
	"time"

	"github.com/TiagoDiass/fastfeet-server/internal/repository"
	"github.com/go-chi/jwtauth"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

type CreateSessionInputDTO struct {
	Document string `json:"document"`
	Password string `json:"password"`
}

type CreateSessionOutputDTO struct {
	AccessToken string `json:"access_token"`
}

type CreateSessionUsecase struct {
	UserRepository repository.UserRepository
	Jwt            *jwtauth.JWTAuth
	JwtExpiresIn   int
}

func NewCreateSessionUsecase(
	userRepository repository.UserRepository,
	jwt *jwtauth.JWTAuth,
	jwtExpiresIn int,
) *CreateSessionUsecase {
	return &CreateSessionUsecase{
		UserRepository: userRepository,
		Jwt:            jwt,
		JwtExpiresIn:   jwtExpiresIn,
	}
}

func (u *CreateSessionUsecase) Execute(input CreateSessionInputDTO) (*CreateSessionOutputDTO, error) {
	user, err := u.UserRepository.FindByDocument(input.Document)

	if err != nil {
		return nil, err
	}

	passwordsMatch := user.ValidatePassword(input.Password)

	if !passwordsMatch {
		return nil, ErrUnauthorized
	}

	_, tokenString, _ := u.Jwt.Encode(map[string]interface{}{
		"sub":  user.ID,
		"exp":  time.Now().Add(time.Second * time.Duration(u.JwtExpiresIn)).Unix(),
		"user": user,
	})

	output := &CreateSessionOutputDTO{
		AccessToken: tokenString,
	}

	return output, nil
}
