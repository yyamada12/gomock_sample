package handler

import (
	"fmt"
	"gin_sample/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestUserHandler_GetUser(t *testing.T) {

	tests := []struct {
		name          string
		reqName       string
		prepareMockFn func(s *service.MockUserService)
		expect        string
	}{
		{
			name:    "test",
			reqName: "test",
			prepareMockFn: func(s *service.MockUserService) {
				s.EXPECT().GetUser("test").Return("testValue", true)
			},
			expect: `{"user":"test","value":"testValue"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// -- arrange --

			// prepare mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserService := service.NewMockUserService(ctrl)
			tt.prepareMockFn(mockUserService)

			h := UserHandler{
				service: mockUserService,
			}

			// prepare gin test context
			w := httptest.NewRecorder()
			c, engine := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("GET", fmt.Sprintf("/user/%s", tt.reqName), nil)
			c.Params = append(c.Params, gin.Param{Key: "name", Value: tt.reqName})

			engine.GET("/user/:name", h.GetUser)

			// -- act --
			h.GetUser(c)
			// ↓のように書くと、Paramsが正しく設定されない
			// engine.HandleContext(c)

			assert.Equal(t, tt.expect, w.Body.String())
		})
	}
}
