package api

import (
	"github.com/gin-gonic/gin"
	"publish-subscribe/service"
	"publish-subscribe/tool"
)

var subs []string
var pubSub = make(map[string][]string)

//创建频道
func createPublisher(ctx *gin.Context) {
	//频道名称
	publisher := ctx.PostForm("publisher")
	_, ok := pubSub[publisher]
	if ok {
		tool.Failure(ctx, 200, "该频道已被创建")
		return
	}
	pubSub[publisher] = subs
	tool.Success(ctx, 200, "你已经成功创建频道"+publisher)
}

//删除频道
func deletePublisher(ctx *gin.Context) {
	publisher := ctx.PostForm("publisher")
	_, ok := pubSub[publisher]
	if !ok {
		tool.Failure(ctx, 200, "该频道不存在")
		return
	}
	delete(pubSub, publisher)
	tool.Success(ctx, 200, "您已成功删除频道"+publisher)
}

//订阅频道
func subscribe(ctx *gin.Context) {
	subscriber := ctx.PostForm("subscriber")
	publisher := ctx.PostForm("publisher")
	_, ok := pubSub[publisher]
	if !ok {
		tool.Failure(ctx, 200, "该频道不存在")
		return
	}
	subs = append(subs, subscriber)
	pubSub[publisher] = subs
	tool.Success(ctx, 200, "您已成功订阅该频道")
}

//发布消息
func publish(ctx *gin.Context) {
	publisher := ctx.PostForm("publisher")
	info := ctx.PostForm("info")
	_, ok := pubSub[publisher]
	if !ok {
		tool.Failure(ctx, 200, "该频道不存在")
		return
	}
	service.Publish(publisher, info, pubSub)
}
