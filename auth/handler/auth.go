package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"

	apierror "github.com/suksest/commodity/lib/error"
	"github.com/suksest/commodity/model/auth"
)

func validateReq(req interface{}) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

func extractToken(req *http.Request) string {
	bearerToken := req.Header.Get("Authorization")
	arr := strings.Split(bearerToken, " ")
	if len(arr) == 2 {
		return arr[1]
	}
	return ""
}

func verifyToken(tokenString string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

// Login will return JWT as response
func Login(c echo.Context) (err error) {
	var loginData auth.LoginRequest
	if err = c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, apierror.ErrResponse{
			Message: apierror.ErrEmptyRequestBody.Error(),
		})
	}

	err = validateReq(loginData)
	if err != nil {
		msg := ""
		field := strings.Trim(strings.Split(err.Error(), " ")[1], "'")
		switch field {
		case "LoginRequest.Phone":
			msg = apierror.ErrIdentifierEmpty.Error()
		case "LoginRequest.Password":
			msg = apierror.ErrPasswordEmpty.Error()
		default:
			msg = apierror.ErrBadInput.Error()
		}
		return c.JSON(http.StatusBadRequest, apierror.ErrResponse{
			Message: msg,
		})
	}

	rsp, err := auth.Login(loginData)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, apierror.ErrResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, rsp)
}

//Signup will return login information
func Signup(c echo.Context) (err error) {
	var signupData auth.SignupRequest
	if err = c.Bind(&signupData); err != nil {
		return c.JSON(http.StatusBadRequest, apierror.ErrResponse{
			Message: apierror.ErrEmptyRequestBody.Error(),
		})
	}

	err = validateReq(signupData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, apierror.ErrResponse{
			Message: apierror.ErrBadInput.Error(),
		})
	}

	rsp, err := auth.Signup(signupData)
	if err != nil {
		if err == apierror.ErrPhoneNotUnique {
			return c.JSON(http.StatusBadRequest, apierror.ErrResponse{
				Message: apierror.ErrPhoneNotUnique.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, apierror.ErrResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, rsp)
}

// Check provided JWT token either valid or not
func Check(c echo.Context) (err error) {
	req := c.Request()
	tokenString := extractToken(req)
	token, err := verifyToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, apierror.ErrResponse{
			Message: apierror.ErrInvalidToken.Error(),
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		name, ok := claims["name"].(string)
		if !ok {
			return err
		}
		phone, ok := claims["phone"].(string)
		if !ok {
			return err
		}
		role, ok := claims["role"].(string)
		if !ok {
			return err
		}
		timestamp, ok := claims["timestamp"].(float64)
		if !ok {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"name":      name,
			"phone":     phone,
			"role":      role,
			"timestamp": time.Unix(int64(timestamp), 0),
		})
	}

	return c.JSON(http.StatusUnauthorized, apierror.ErrResponse{
		Message: apierror.ErrInvalidToken.Error(),
	})
}
