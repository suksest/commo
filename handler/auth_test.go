package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	testIn := []struct {
		desc    string
		reqBody string
		resBody string
		resCode int
	}{
		{
			desc:    "With empty body",
			reqBody: ``,
			resBody: `{"message":"request.empty_request_body"}`,
			resCode: http.StatusBadRequest,
		},
		{
			desc:    "With empty identifier",
			reqBody: `{"phone": "", "password":"test"}`,
			resBody: `{"message":"auth.identifier_empty"}`,
			resCode: http.StatusBadRequest,
		},
		{
			desc:    "With empty password",
			reqBody: `{"phone": "test", "password": ""}`,
			resBody: `{"message":"auth.password_empty"}`,
			resCode: http.StatusBadRequest,
		},
	}

	for _, in := range testIn {
		e := echo.New()
		t.Log(in.desc)

		req, err := http.NewRequest(
			echo.POST,
			"/auth/login",
			strings.NewReader(in.reqBody),
		)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		rsp := httptest.NewRecorder()

		c := e.NewContext(req, rsp)

		if assert.NoError(t, Login(c)) {
			assert.Equal(t, in.resCode, rsp.Code)
			assert.Equal(t, in.resBody, strings.Trim(rsp.Body.String(), "\n"))
		}
	}
}

func TestSignup(t *testing.T) {
	testIn := []struct {
		desc    string
		reqBody string
		resBody string
		resCode int
	}{
		{
			desc:    "With empty body",
			reqBody: ``,
			resBody: `{"message":"request.empty_request_body"}`,
			resCode: http.StatusBadRequest,
		},
		{
			desc:    "With empty phone",
			reqBody: `{"phone": "", "name":"test", "role":"user"}`,
			resBody: `{"message":"request.input_invalid"}`,
			resCode: http.StatusBadRequest,
		},
		{
			desc:    "With empty password",
			reqBody: `{"phone": "", "name":"test", "role":"user"}`,
			resBody: `{"message":"request.input_invalid"}`,
			resCode: http.StatusBadRequest,
		},
		{
			desc:    "With empty role",
			reqBody: `{"phone": "", "name":"test", "role":"user"}`,
			resBody: `{"message":"request.input_invalid"}`,
			resCode: http.StatusBadRequest,
		},
	}

	for _, in := range testIn {
		e := echo.New()
		t.Log(in.desc)

		req, err := http.NewRequest(
			echo.POST,
			"/auth/signup",
			strings.NewReader(in.reqBody),
		)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set(
			echo.HeaderContentType,
			echo.MIMEApplicationJSON,
		)

		rsp := httptest.NewRecorder()

		c := e.NewContext(req, rsp)

		if assert.NoError(t, Signup(c)) {
			assert.Equal(t, in.resCode, rsp.Code)
			assert.Equal(t, in.resBody, strings.Trim(rsp.Body.String(), "\n"))
		}
	}
}
