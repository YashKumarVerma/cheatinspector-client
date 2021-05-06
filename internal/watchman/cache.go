package watchman

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/YashKumarVerma/hentry-client/internal/fs"
	"github.com/allegro/bigcache/v3"
	"time"
)


var cache *bigcache.BigCache
var network bytes.Buffer
var encoder *gob.Encoder
var decoder *gob.Decoder

// Init to setup and initialize cache handlers
func Init() {
	// initialize cache handler
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	cache.Set("status", []byte("Cache Serve Alive"))
	entry, _ := cache.Get("status")
	fmt.Println(string(entry))

	// initialise encoders and decoders
	encoder = gob.NewEncoder(&network)
	decoder = gob.NewDecoder(&network)
}

// setCache saves the key and vale into storage
func setCache(fileDetails fs.FileDetails) bool {
	err := encoder.Encode(fileDetails)
	if err != nil {
		fmt.Println("Error encoding value : " , fileDetails)
		cache.Set(fileDetails.Path, network.Bytes())
	}
	//cache.Set(index, []byte(fileDetails))
}

