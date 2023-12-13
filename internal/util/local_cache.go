package util

import (
	"sync"
	"time"
)

// LocalCache 一个简单的本地永久缓存,数据过期时触发更新数据
type LocalCache struct {
	Value      interface{}
	Expiration time.Time
	Duration   time.Duration
	DataFunc   func() interface{}
	mu         sync.RWMutex
}

func NewLocalCache(dataFunc func() interface{}, duration time.Duration) *LocalCache {
	cache := &LocalCache{
		Duration: duration,
		DataFunc: dataFunc,
	}
	cache.update()
	return cache
}
func (c *LocalCache) update() {
	if c.mu.TryLock() { // 只允许一个协程触发更新拉取
		c.Value = c.DataFunc()
		c.Expiration = time.Now().Add(c.Duration)
		defer c.mu.Unlock()
	}
}
func (c *LocalCache) Get() interface{} {
	// 如果超过访问次数阈值或过期时间，则重新拉取数据
	if c.Value == nil || c.Expiration.Before(time.Now()) {
		go c.update()
	}
	return c.Value
}
