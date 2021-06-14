package helpers

import (
	"github.com/pmylund/go-cache"
	"log"
	"time"
)

type CacheService interface {
	GetCache(key string) (interface{}, bool)
	SetCache(key string, response interface{})
}

type cacheServer struct {
	cacheServer *cache.Cache
}

func NewCacheServer() *cacheServer {
	return &cacheServer{
		ck,
	}
}

//using go-cache for caching, but can be easily swapped with redis/memcached etc
var ck = cache.New(24*time.Hour, 12*time.Hour)

func (c *cacheServer) GetCache(ifsc string) (interface{}, bool) {
	return c.cacheServer.Get(ifsc)
}

func (c *cacheServer) SetCache(ifsc string, response interface{}) {
	log.Println("cache set")
	c.cacheServer.Set(ifsc, response, 0)
}
