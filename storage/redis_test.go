package storage

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"testing"
)

func TestUnit(t *testing.T) {
	rdb := NewClient("192.168.1.5:6379")
	err := rdb.Set("key2", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("key2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key2 = ", val)
}
func TestExample(t *testing.T) {
	//var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
