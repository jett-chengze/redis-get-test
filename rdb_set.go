package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_set(rdb *redis.Client, ctx context.Context) {
	log.Println()
	resSet1_Set1_Val0IsMember, err := rdb.SIsMember(ctx, "set1", "set1_Val0").Result()
	if err != nil {
		log.Fatalf("redis sismember set1 error: %s", err)
	}
	log.Printf("redis sismember set1 set1_Val0: %v", resSet1_Set1_Val0IsMember)

	resSet1, err := rdb.SMembers(ctx, "set1").Result()
	if err != nil {
		log.Fatalf("redis smembers set1 error: %s", err)
	}
	log.Printf("redis smembers set1: %v", resSet1)
}
