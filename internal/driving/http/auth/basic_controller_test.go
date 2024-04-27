package auth

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ricardo134/auth-service/internal/core/app/auth"
	"testing"
)

func Test_basicController_Login(t *testing.T) {
	type fields struct {
		authr auth.AuthenticateService
	}
	type args struct {
		gtx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := basicController{
				authr: tt.fields.authr,
			}
			c.Login(tt.args.gtx)
		})
	}
}
