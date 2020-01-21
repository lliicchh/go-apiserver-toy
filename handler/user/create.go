package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	. "github.com/lliicchh/apiserver/handler"
	"github.com/lliicchh/apiserver/model"
	"github.com/lliicchh/apiserver/pkg/errno"
	"github.com/lliicchh/apiserver/util"
)

func Get(c *gin.Context) {

}
func List(c *gin.Context) {

}
func Delete(c *gin.Context) {

}
func Update(c *gin.Context) {

}
func Create(c *gin.Context) {

	// version 3
	log.Infof("user create function called", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// 参数校验
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
	}

	// 密码加密
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
	}

	// 创建用户，往数据库添加记录
	if err := u.Create(); err != nil {
		log.Debugf("err : %s", err)
		SendResponse(c, errno.ErrDatabase, nil)
	}

	log.Info("inser ok")
	// 结果返回
	resp := CreateResponse{Username: r.Username}
	SendResponse(c, nil, resp)
}
