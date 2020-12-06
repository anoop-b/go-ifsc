package helpers

import (
	"log"
	"time"

	"github.com/pmylund/go-cache"
)

//using go-cache for caching, but can be easily swapped with redis/memcached etc
var ck = cache.New(24*time.Hour, 12*time.Hour)

func GetCache(ifsc string) (interface{}, bool) {
	return ck.Get(ifsc)
}

func SetCache(ifsc string, response interface{}) {
	log.Println("cache set")
	ck.Set(ifsc, response, 0)
}
