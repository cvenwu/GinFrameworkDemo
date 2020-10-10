package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/10/10 6:14 下午
 * @Desc: json格式数据进行返回：
			1. 使用map作为json格式进行返回
			2. 使用struct类型作为json格式进行返回
 */

func main() {
	engine := gin.Default()

	//使用map作为json格式的数据进行返回
	engine.GET("/mapjson", func(context *gin.Context) {
		fullpath := "请求路径：" + context.FullPath()

		context.JSON(http.StatusOK, map[string]interface{}{
			"code":    1,
			"message": "OK",
			"data":    fullpath,
		})

	})

	//使用struct作为json格式的数据进行返回
	engine.GET("/structjson", func(context *gin.Context) {
		fullpath := context.FullPath()

		resp := Response{
			Code:    1,
			Message: "OK",
			Data:    fullpath,
		}
		context.JSON(http.StatusOK, &resp)
	})

	engine.Run()

}

//因为我们要返回json格式的数据所以我们需要提前定义一个结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
