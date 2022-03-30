package main

import (
	"publish-subscribe/api"
	"publish-subscribe/dao"
)

func main() {
	dao.InitRedis()
	api.InitEngine()
}
