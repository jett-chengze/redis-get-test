package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_hash(rdb *redis.Client, ctx context.Context) {
	log.Println()
	resHash1Field1, err := rdb.HGet(ctx, "hash1", "field1").Result()
	if err != nil {
		log.Fatalf("redis hget hash1 error: %s", err)
	}
	log.Printf("redis hget hash1 field1: %s", resHash1Field1)

	resHash1Field0Exists, err := rdb.HExists(ctx, "hash1", "field0").Result()
	if err != nil {
		log.Fatalf("redis hexists hash1 field0 error: %s", err)
	}
	log.Printf("redis hexists hash1 field0: %v", resHash1Field0Exists) //true or false

	resHash1, err := rdb.HGetAll(ctx, "hash1").Result()
	if err != nil {
		log.Fatalf("redis hgetall hash1 error: %s", err)
	}
	log.Printf("redis hgetall hash1: %v", resHash1) //map[string]string  map[key1:value1 key2:value2 ...]
}
