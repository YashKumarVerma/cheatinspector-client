package watchman

import (
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func Init() {
	cache,_ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	cache.Set("status", []byte("Cache Serve Alive"))
	entry, _ := cache.Get("status")
	fmt.Println(string(entry))

}

var cache *bigcache.BigCache
