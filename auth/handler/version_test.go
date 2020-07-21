package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	testIn := []struct {
		desc    string
		resBody string
		resCode int
	}{
		{
			desc:    "Truthy scenarion",
			resBody: `{"backend":"v0"}`,
			resCode: http.StatusOK,
		},
	}
	for _, in := range testIn {
		e := echo.New()
		t.Log(in.desc)

		req, err := http.NewRequest(
			echo.GET,
			"/version",
			strings.NewReader(""),
		)
		if err != nil {
			t.Fatal(err)
		}

		rsp := httptest.NewRecorder()

		c := e.NewContext(req, rsp)

		if assert.NoError(t, GetVersion(c)) {
			assert.Equal(t, in.resCode, rsp.Code)
			assert.Equal(t, in.resBody, strings.Trim(rsp.Body.String(), "\n"))
		}
	}
}
