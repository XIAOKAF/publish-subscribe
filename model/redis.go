package model

type Publisher struct {
	Publisher   string
	RedisServer RedisServer
}

type RedisServer struct {
	PubSubChannel map[string][]string //保存订阅频道的状态（频道名称与关注该频道的订阅者）
}
