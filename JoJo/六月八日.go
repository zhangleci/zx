package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//第二课作业
	r.GET("/name", func(c *gin.Context) {
		//nickName 接口传过来的昵称
		nickName := c.Query("nickName")
		//result 存放计算出来的真实名
		result := ""

		switch nickName {
		case "愿":
			result = "聂昶如"
		case "JoJo":
			result = "张涵"
		case "Re":
			result = "王苗"
		case "Au":
			result = "学长"
		case "Z":
			result = "张鑫"
		case "Qiu":
			result = "马伟成"
		default:
			result = "can not find this people."
		}
		c.JSON(200, gin.H{
			"realName": result,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
