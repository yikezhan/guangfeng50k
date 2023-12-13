package redisTemplate

import "fmt"

const (
	pre              = "ET:"
	HotGoodsIdSetKey = pre + "hot:g:ids"
	GoodsInfo        = pre + "goods:id:%v"
	UserIdInfo       = pre + "user:id:%v:password:%v"
)

func BuildKey(key string, v ...any) string {
	return fmt.Sprintf(key, v...)
}
