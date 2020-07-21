package auth

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	apierror "github.com/suksest/commodity/lib/error"

	"github.com/suksest/commodity/model"
)

// LoginRequest represent request for login endpoint
type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse represent response for login endpoint
type LoginResponse struct {
	Token string `json:"token"`
}

// SignupRequest represent request for signup endpoint
type SignupRequest struct {
	Phone string `json:"phone" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Role  string `json:"role" validate:"required"`
}

// SignupResponse represent response for signup endpoint
type SignupResponse struct {
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

// AccessDetails represent response for authentication checker
type AccessDetails struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Timestamp time.Time `json:"timestamp"`
	ExpireAt  time.Time `json:"expire_at"`
}

func compareHashAndPassword(userPassword, passwordIn string) error {
	h := md5.New()
	io.WriteString(h, passwordIn)

	s := fmt.Sprintf("%x", h.Sum(nil))

	if strings.Compare(s, userPassword) != 0 {
		return apierror.ErrPasswordWrong
	}

	return nil
}

// Login will check identifier password pair
func Login(loginData LoginRequest) (rsp *LoginResponse, err error) {
	user, err := model.FindUserByPhone(loginData.Phone)
	if err != nil {
		return nil, err
	}

	err = compareHashAndPassword(user.Password, loginData.Password)
	if err != nil {
		return nil, err
	}

	payload := jwt.MapClaims{}
	payload["name"] = user.Name
	payload["phone"] = user.Phone
	payload["role"] = user.Role
	payload["timestamp"] = time.Now().Unix()
	payload["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := NewToken(payload)
	jwtToken, err := GetSignedToken(token)

	if err != nil {
		return nil, err
	}

	rsp = &LoginResponse{
		Token: jwtToken,
	}

	return
}

// generatePassword will generate 4(four) random alphanumeric characters
func generatePassword() (password string) {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 4
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	password = b.String() // E.g. "e8Yr"
	return
}

// Signup will check provided sign up request
// return JWT auth token
func Signup(signupData SignupRequest) (rsp *SignupResponse, err error) {
	user, _ := model.FindUserByPhone(signupData.Phone)
	// check if there is a user with provided phone
	if user != nil {
		return nil, apierror.ErrPhoneNotUnique
	}

	req := model.User{
		Name:     signupData.Name,
		Phone:    signupData.Phone,
		Role:     signupData.Role,
		Password: generatePassword(),
	}

	res, err := model.CreateUser(req)
	if err != nil {
		return nil, err
	}
	rsp = &SignupResponse{
		Name:     res.Name,
		Phone:    res.Phone,
		Role:     res.Role,
		Password: req.Password,
	}

	return
}
