package main

import (
	"fmt"
	"net/url"
)

func login(ctx*gin.Context) {
	username := ctx.PostForm ("username")
	password := ctx.PostForm ( "password ")
 flag,err := sevice.IsPasswordCorrect (username ,password)
 if err != nil {
	 fmt .Print ("judge password correct err:", err)

	 tool.RespInternalError(ctx)
	 return
 }
 if !flag {
	 tool.RespErrorWithDate (ctx,date:"密码错误")
	 return
 }

ctx.SetCookie (name :"username",username ,maxAge :100,path"/" ,domain"" ,secure:false,httpOnly:false)
 tool.RespSuccessful(ctx)



}
