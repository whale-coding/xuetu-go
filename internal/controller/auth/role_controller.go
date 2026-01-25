package auth

import (
	"xuetu-project/internal/service"

	"github.com/gin-gonic/gin"
)

type RoleController interface {
	getRole(ctx *gin.Context)
}

type roleController struct {
	svc *service.Service
}

func NewRoleController(svc *service.Service) RoleController {
	return &roleController{svc: svc}
}

func (c *roleController) getRole(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
