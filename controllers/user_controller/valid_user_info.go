package user_controller

import (
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseInfo struct {
	Email    string `form:"email" json:"email" validate:"required"`
	Username string `form:"username" json:"username" validate:"required"`
}

// validate base info of user when register
func ValidBaseInfo(ctx *gin.Context) {
	res := helper.Res{}

	var baseInfo BaseInfo
	if err := ctx.Bind(&baseInfo); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	// user does exist
	if _, err := models.FindOneByUsername(baseInfo.Username); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	// the email of user does exist
	if _, err := models.FindOneByEmail(baseInfo.Email); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	res.Success(gin.H{
		"isValid": true,
	}).Send(ctx)
}
