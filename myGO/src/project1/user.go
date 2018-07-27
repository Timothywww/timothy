package handler

import (
	"project1_model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

func (h *Handler) Signup(ctx echo.Context) (err error) {
	user := model.User{}
	user.Group = 0
	if err = ctx.Bind(&user); err != nil {
		return ctx.JSON(401, "提交的格式不对")
	}

	db := h.DB.Clone()
	defer db.Close()

	c := db.DB("test").C("users")

	if err = c.Find(bson.M{"account":user.Account}).One(&user);err == nil {
		return 
	}

	if err = c.Insert(&user); err != nil {
		return ctx.JSON(403, "没有找到")
	}

	return ctx.JSON(200, user)

}

//Log in...
func (h *Handler) Login(ctx echo.Context) (err error) {
	user := model.User{}
	if err = ctx.Bind(&user); err != nil {
		return ctx.JSON(401, "提交的格式不对")
	}

	db := h.DB.Clone()
	defer db.Close()
	c := db.DB("test").C("users")

	if err = c.Find(bson.M{"account": user.Account, "password": user.Password}).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			// return &echo.HTTPError {Code: http.StatusUnauthorized, Message:"user not found"}
			return ctx.JSON(403, "没有找到")
		}
		return ctx.JSON(500, err)
	}
	if user.Group == 1{
		return ctx.JSON(200,"Welcome, admin")
	}
	return ctx.JSON(200, "Succeed")
}
