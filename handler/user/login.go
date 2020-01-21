package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/lliicchh/apiserver/handler"
	"github.com/lliicchh/apiserver/model"
	"github.com/lliicchh/apiserver/pkg/auth"
	"github.com/lliicchh/apiserver/pkg/errno"
	"github.com/lliicchh/apiserver/pkg/token"
)

func Login(c *gin.Context) {
	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Username)

	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPassWordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	SendResponse(c, nil, model.Token{Token: t})

}
