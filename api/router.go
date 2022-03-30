package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()
	publisherGroup := engine.Group("/publisher")
	{
		publisherGroup.POST("/createPublisher", createPublisher) //创建频道
		publisherGroup.POST("/deletePublisher", deletePublisher) //删除频道
		//publisherGroup.POST("/publish", publish)                 //发布消息
	}
	//engine.POST("/subscribe", subscribe) //订阅频道
	engine.Run(":8080")
}
