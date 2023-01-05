package controller

import (
	"fmt"
	"goblog-end/dao"
	"goblog-end/model"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	// 获取路径参数
	// name := c.Param("name")

	var dataList []model.BlogPost

	//查询数据库
	dataList = dao.GetAllPost()
	for i := 0; i < len(dataList); i++ {
		if len(dataList[i].Content) > 100 {
			dataList[i].Content = dataList[i].Content[0:100] + "..."
		}
		fmt.Printf("v: %v\n", dataList[i].Content)

	}

	//判断是否查到数据
	if len(dataList) == 0 {
		c.JSON(200, gin.H{
			"meg":  "没有查到数据",
			"code": 400,
			"data": gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"meg":  "查询成功",
			"code": 200,
			"data": dataList,
		})
	}

}

func PostByPid(c *gin.Context) {
	// 获取路径参数
	pid := c.Param("pid")

	var dataList []model.BlogPost

	//查询数据库
	dataList = dao.GetAllPost()
	for i := 0; i < len(dataList); i++ {
		if len(dataList[i].Content) > 100 {
			dataList[i].Content = dataList[i].Content[0:100] + "..."
		}
		fmt.Printf("v: %v\n", dataList[i].Content)

	}

	//判断是否查到数据
	if len(dataList) == 0 {
		c.JSON(200, gin.H{
			"meg":  "没有查到数据",
			"code": 400,
			"data": gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"meg":  "查询成功",
			"code": 200,
			"data": dataList,
		})
	}

}
