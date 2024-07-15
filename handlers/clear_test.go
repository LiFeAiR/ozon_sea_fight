package handlers

import (
	"github.com/LiFeAiR/ozon_sea_fight/dto"
	"github.com/LiFeAiR/ozon_sea_fight/handlers/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClear(t *testing.T) {
	type args struct {
		panic bool
		dto   dto.State
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody string
	}{
		{"success", args{false, dto.State{
			ShipCount:    1,
			Destroyed:    2,
			Knocked:      3,
			KnockedSlice: []int{4, 5},
			ShotCount:    6,
		}}, 200, `{"ship_count":1,"destroyed":2,"knocked":3,"shot_count":6}`},
		{"error", args{true, dto.State{}}, 500, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			router := gin.Default()
			a := &mocks.AppMock{}
			if tt.args.panic {
				a.On("Clear").Panic("test: panic recovered")
			} else {
				a.On("Clear").Return()
			}
			a.On("State").Return(tt.args.dto)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/test", nil)

			// Act
			router.POST("/test", Clear(a))
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tt.wantCode, w.Code)
			assert.Equal(t, w.Body.String(), tt.wantBody)
		})
	}
}
