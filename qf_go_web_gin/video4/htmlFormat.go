package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/10/10 6:22 下午
 * @Desc: 返回html数据
 */

func main() {
	engine := gin.Default()

	//gin如果要返回静态资源，需要先进行设置
	engine.LoadHTMLGlob("/Users/yirufeng/实习/项目/GinFrameworkDemo/templates/*")
	//engine.LoadHTMLGlob("./templates/*") 自己使用这个会报错，所以就使用绝对路径进行代替

	//如果要引入图片等静态资源文件，需要先进行设置
	//第1个参数是页面加载图片对应的src的路径，也就是页面中的图片发现有/img的时候会去我们第2个参数里面的目录找
	//这里实际上就是一个映射关系，将我们src里面用到的/img映射到路径./imgs
	//engine.Static("/img", "/Users/yirufeng/实习/项目/GinFrameworkDemo/img")
	//下面这个也会起作用的
	engine.Static("/img", "../../img")

	engine.GET("/index", func(context *gin.Context) {
		//获取客户端请求路径
		fullpath := context.FullPath()
		log.Println(fullpath)

		//将html文件进行返回
		//第1个是状态码，第2个是返回的模板的文件对应的文件名，第3个是模板中的参数
		//注意：gin.H不仅可以用于作为JSON返回的时候的data还可以给模板填充参数
		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullpath": fullpath,
			"title": "gin教程",
		})
	})

	engine.Run()
}
