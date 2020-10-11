package controller

import "github.com/gin-gonic/gin"

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/10/11 11:55 上午
 * @Desc:
 */

type HelloController struct {

}

//如何让engine解析/hello呢
//参数中得到的engine与我们的main.go中的app是对应的
func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.Hello)

}

//解析 /hello请求
func (hello *HelloController) Hello(context *gin.Context) {
	context.Writer.Write([]byte("hello,"))
}