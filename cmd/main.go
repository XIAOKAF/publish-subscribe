package cmd

import (
	"publish-subscribe/dao"
)

func main() {
	dao.InitRedis()
}
