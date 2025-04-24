package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_list(rdb *redis.Client, ctx context.Context) {
	log.Println()
	resList1, err := rdb.LRange(ctx, "list1", 0, 2).Result()
	if err != nil {
		log.Fatalf("redis lrange error: %s", err)
	}
	log.Printf("redis lrange list1: %v", resList1) //[value1 value2 ...]
}
