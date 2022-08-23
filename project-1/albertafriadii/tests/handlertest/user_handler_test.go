package handlertest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	samples := []struct {
		inputJSON    string
		statusCode   int
		username     string
		errorMessage string
	}{
		{
			inputJSON:    `{"username": "frank", "email": "frank@gmail.com", "password_1": "Password00_", "password_2": "Password00_"}`,
			statusCode:   201,
			username:     "frank",
			errorMessage: "",
		},
		{
			inputJSON:    `{"username": "", "email": "frank@gmail.com", "password_1": "Password00_", "password_2": "Password00_"}`,
			statusCode:   400,
			errorMessage: "Required Username",
		},
		{
			inputJSON:    `{"username": "Ken", "email": "", "password_1": "Password00_", "password_2": "Password00_"}`,
			statusCode:   400,
			errorMessage: "Required Email",
		},
		{
			inputJSON:    `{"username": "frank", "email": "frank@gmail.com", "password_1": "", "password_2": ""}`,
			statusCode:   400,
			errorMessage: "Required Password",
		},
		{
			inputJSON:    `{"username": "Ken", "email": "ken@gmail.com", "password_1": "Password00_", "password_2": "Password00_1"}`,
			statusCode:   400,
			errorMessage: "Password not match",
		},
		{
			inputJSON:    `{"username": "Ken", "email": "ken@gmail.com", "password_1": "password", "password_2": "password"}`,
			statusCode:   400,
			errorMessage: "Must be at least 8 letters, alphanumeric + symbol, has at least 1 uppercase letter, has at least 1 number, and has at least 1 symbol.",
		},
	}

	for _, v := range samples {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewBufferString(v.inputJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		h := echo.HandlerFunc(server.CreateUser)
		c := e.NewContext(req, rec)

		if assert.NoError(t, h(c)) {
			responseMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rec.Body.String()), &responseMap)
			if err != nil {
				fmt.Printf("Cannot convert to json: %v", err)
			}
			assert.Equal(t, rec.Code, v.statusCode)
			if v.statusCode == 201 {
				assert.Equal(t, responseMap["username"], v.username)
			}
			if v.statusCode == 400 {
				assert.Equal(t, responseMap["error"], v.errorMessage)
			}
		}
	}
}

// func TestGetUser(t *testing.T) {
// 	err := refreshUserTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	user, err := oneUser()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	userSample := []struct {
// 		user_id      string
// 		statusCode   int
// 		username     string
// 		errorMessage string
// 	}{
// 		{
// 			user_id:    user.UserId,
// 			statusCode: 200,
// 			username:   user.Username,
// 		},
// 		{
// 			user_id:    "unknown",
// 			statusCode: 400,
// 		},
// 	}

// 	for _, v := range userSample {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodGet, "/user/get", nil)
// 		rec := httptest.NewRecorder()
// 		h := echo.HandlerFunc(server.CreateUser)
// 		c := e.NewContext(req, rec)
// 	}
// }
