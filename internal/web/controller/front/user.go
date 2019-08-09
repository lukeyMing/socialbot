package front

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialbot/internal/web/common"
	"socialbot/internal/web/logic/userLogic"
	"socialbot/internal/web/model"
	"socialbot/pkg/utils"
)

func Login(c *gin.Context) {
	var accountForm model.AccountForm
	if err := c.ShouldBind(&accountForm); err != nil {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}

	result := userLogic.Login(accountForm)
	c.JSON(http.StatusOK, result)
}

func Register(c *gin.Context) {
	var accountForm model.AccountForm
	if err := c.ShouldBind(&accountForm); err != nil {
		c.JSON(http.StatusOK, common.ParamError)
		return
	}
	result := userLogic.Register(accountForm, c.ClientIP())
	c.JSON(http.StatusOK, result)
}

func EditProfile(c *gin.Context)  {
	nickname := utils.TrimSpace(c.DefaultPostForm("nickname", ""))
	intro := utils.TrimSpace(c.DefaultPostForm("intro", ""))
	if nickname == "" {
		c.JSON(http.StatusOK, common.NicknameEmpty)
		return
	}
	userLogic.EditProfile(c, nickname, intro).Out(c)
}