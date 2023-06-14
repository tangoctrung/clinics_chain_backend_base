package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/minh1611/clinics_chain_management/authservice/src/internal/lib"
	"github.com/minh1611/clinics_chain_management/authservice/src/pb/authpb"
)

var UserSet = wire.NewSet(wire.Struct(new(UserController), "*"), wire.Struct(new(UserService), "*"))

type UserController struct {
	S *UserService
}

func (s *UserController) HandleGet(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (s *UserController) HandlePost(g *gin.Context) {
	req := authpb.NewUser{
		Name: g.GetString("name"),
		//Age:         g.GetString("age"),
		PhoneNumber: g.GetString("phoneNumber"),
		Email:       g.GetString("email"),
		Password:    g.GetString("password"),
	}
	if err := lib.GetBody(g, &req); err != nil {
		lib.BadRequest(g, err)
		return
	}
	res := authpb.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
		Id:          1,
	}
	lib.Success(g, &res)
}
