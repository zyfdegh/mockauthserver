package main

import (
	"log"

	"github.com/kataras/iris"
)

// Resp is the common struct of http response body
type Resp struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

func main() {

	iris.Post("/user/login", handleRoot)

	iris.Listen(":8080")
}

func handleRoot(ctx *iris.Context) {
	email := ctx.FormValueString("email")
	password := ctx.FormValueString("password")

	resp := Resp{}
	ok, err := auth(email, password)
	if err != nil {
		log.Printf("auth failed: %v\n", err)
		resp.Code = iris.StatusInternalServerError
		resp.Msg = "auth error"
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}
	if !ok {
		resp.Code = 401
		resp.Msg = "not ok"
		ctx.JSON(iris.StatusUnauthorized, resp)
		return
	}
	resp.Code = 200
	resp.Msg = "ok"
	resp.Token = "1234567890"
	ctx.JSON(iris.StatusOK, resp)
	return
}

func auth(email, password string) (bool, error) {
	if email == "admin@email.com" && password == "linker" {
		return true, nil
	}
	return false, nil
}
