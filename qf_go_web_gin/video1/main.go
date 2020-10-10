package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/10/10 9:15 上午
 * @Desc: 1
 */

func main() {
	engine := gin.Default()

	engine.GET("/", func(context *gin.Context) {
		log.Println(context.FullPath())
		context.Writer.Write([]byte("hello world"))
	})

	if err := engine.Run(); nil != err {
		log.Fatal(err)
	}
}