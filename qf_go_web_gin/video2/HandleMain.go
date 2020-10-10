package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/10/10 9:36 上午
 * @Desc: 2-1 使用handle的方式对请求进行处理
			gin中使用Handle处理http请求
 */

func main() {
	engine := gin.Default()


	//模仿一个get请求
	//context是gin框架帮助我们封装好的方便我们进行Gin框架操作的资源和对象的结构体对象
	//我们一般叫做上下文环境变量
	engine.Handle(http.MethodGet, "/hello", func(context *gin.Context) {
		log.Println(context.FullPath())  //获取本次请求的接口

		//从url中获取参数，第1个为Key第2个为默认值
		name := context.DefaultQuery("name", "yirufeng")
		fmt.Println(name)

		//Write表示向浏览器将数据返回给前端
		context.Writer.Write([]byte("hello," +  name))
	})

	//模仿一个Post请求，用户要进行登录
	//post请求会将数据放在请求body中
	engine.Handle(http.MethodPost, "/login", func(context *gin.Context) {
		log.Println(context.FullPath())  //获取本次请求的接口

		//使用PostForm解析表单中的参数，
		//需要传入一个key，表示我们想要获取post表单中的哪一个字段
		username := context.PostForm("username")
		password := context.PostForm("password")
		log.Println("用户名：", username, "；密码：", password)

		context.Writer.Write([]byte(username + "进行了登录"))
	})

	engine.Run()
}