package redisTemplate

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"time"
)

const (
	_defaultTime = time.Second
)

func Set(rdb *redis.Client, key string, value any) {
	val, _ := json.Marshal(value)
	rdb.Set(context.Background(), key, val, _defaultTime*3600)
}
func Get(rdb *redis.Client, key string) []byte {
	val, _ := rdb.Get(context.Background(), key).Bytes()
	return val
}
func GetSCmd(rdb *redis.Client, key string) *redis.StringCmd {
	return rdb.Get(context.Background(), key)
}
func Del(rdb *redis.Client, key string) {
	rdb.Del(context.Background(), key)
}

func ZRange(rdb *redis.Client, key string, start, stop int64) ([]string, error) {
	return rdb.ZRange(context.Background(), key, start, stop).Result()
}
func ZSet(rdb *redis.Client, key string, hot, goodsId int64) {
	obj := &redis.Z{
		Score:  float64(hot),
		Member: goodsId,
	}
	rdb.ZAdd(context.Background(), key, *obj)
}
func ZRemRangeByScore(rdb *redis.Client, key string, min, max string) {
	rdb.ZRemRangeByScore(context.Background(), key, min, max)
}
