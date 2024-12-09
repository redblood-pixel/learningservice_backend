package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/redblood-pixel/learning-service-go/internal/tokenutil"
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	"github.com/redblood-pixel/learning-service-go/pkg/service"
	mock_service "github.com/redblood-pixel/learning-service-go/pkg/service/mocks"
	"go.uber.org/mock/gomock"
)

func TestHandler_sighUp(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockUsers, input domain.SignupInput)

	testTable := []struct {
		name                string
		inputBody           string
		input               domain.SignupInput
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"chill_guy","email":"lala1@mail.ru","password":"qye1243"}`,
			input: domain.SignupInput{
				Username: "chill_guy",
				Email:    "lala1@mail.ru",
				Password: "qye1243",
			},
			mockBehaviour: func(s *mock_service.MockUsers, input domain.SignupInput) {
				s.EXPECT().SignUp(input)
			},
			expectedStatusCode: 200,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			users := mock_service.NewMockUsers(c)
			testCase.mockBehaviour(users, testCase.input)

			tm := tokenutil.NewTokenManager(&tokenutil.Config{
				AccessTokenTTL:  1 * time.Hour,
				RefreshTokenTTL: 3000 * time.Hour,
				SigningKey:      "12341234",
			})
			services := &service.Service{Users: users}
			handler := NewHandler(services, tm)

			// Test server
			r := gin.New()
			r.POST("/sign-up", handler.userSignup)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			// Perform request
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}
