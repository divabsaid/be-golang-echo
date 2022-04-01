package usecase

import (
	"be-golang-echo/entity/user"
	"be-golang-echo/entity/user/repository"
	"be-golang-echo/utils/jwt"
	"be-golang-echo/utils/password"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserUsecase interface {
	UserRegister(um *user.UserModel) (*user.UserModel, error)
	UserLogin(um *user.UserLoginModel) (map[string]string, error)
	GetProfile(id int) (*user.UserProfileModel, error)
	RequestToken(id int) (map[string]string, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(u repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: u,
	}
}

func (u *userUsecase) UserRegister(um *user.UserModel) (*user.UserModel, error) {
	err := validator.New().Struct(um)

	if err != nil {
		return um, user.REQUEST_BODY_NOT_VALID
	}
	um.Active = true
	um.RoleID = 1
	um.Password, _ = password.HashPassword(um.Password)
	um.CreatedAt = time.Now()
	res, err := u.userRepository.UserRegister(um)
	if err != nil {
		return um, err
	}
	return res, nil
}

func (u *userUsecase) UserLogin(um *user.UserLoginModel) (map[string]string, error) {
	err := validator.New().Struct(um)
	if err != nil {
		return nil, user.REQUEST_BODY_NOT_VALID
	}
	userObj, err := u.userRepository.UserLogin(um)
	if err != nil {
		return nil, err
	}
	if !userObj.Active {
		return nil, user.LOGIN_FAILED_INACTIVE
	}
	loggedIn, err := password.VerifyPassword(um.Password, userObj.Password)
	if err != nil || !loggedIn {
		return nil, user.LOGIN_FAILED
	}
	token, _ := jwt.CreateJWTToken(userObj)

	return token, nil
}

func (u *userUsecase) GetProfile(id int) (*user.UserProfileModel, error) {
	res, err := u.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	profile := new(user.UserProfileModel)
	profile.Email = res.Email
	profile.Username = res.Username
	profile.Fullname = res.Fullname
	profile.RoleName = res.RoleName
	profile.ImageName = res.ImageName

	return profile, nil
}

func (u *userUsecase) RequestToken(id int) (map[string]string, error) {
	res, err := u.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	userObj := new(user.UserLoginModel)
	userObj.ID = res.ID
	userObj.RoleID = res.RoleID

	token, _ := jwt.CreateJWTToken(userObj)

	return map[string]string{
		"access_token": token["access_token"],
	}, nil
}
