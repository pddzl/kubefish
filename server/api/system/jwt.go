package system

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/common/response"
	"github.com/pddzl/kubefish/server/model/system"
	"go.uber.org/zap"
)

type JwtApi struct{}

// JoinInBlacklist
// @Tags      Jwt
// @Summary   jwt加入黑名单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "jwt加入黑名单"
// @Router    /jwt/jsonInBlacklist [post]
func (j *JwtApi) JoinInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	err := jwtService.JoinInBlacklist(jwt)
	if err != nil {
		global.KF_LOG.Error("jwt作废失败", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
		return
	}
	response.OkWithMessage("jwt作废成功", c)
}
