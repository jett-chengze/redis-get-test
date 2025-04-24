package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_hash_ByLua(rdb *redis.Client, ctx context.Context) {
	log.Println()
	luaScript_Hash1Field1 := `
		local r1 = ""
		r1 = redis.call("HGET", "hash1", "field1")
		return r1
	`
	resHash1Field1, err := rdb.Eval(ctx, luaScript_Hash1Field1, []string{}).Result()
	if err != nil {
		log.Fatalf("redis hget hash1 error: %s", err)
	}
	log.Printf("redis hget hash1 field1: %s", resHash1Field1)

	luaScript_Hash1Field0Exist := `
		local r1 = ""
		r1 = redis.call("HEXISTS", "hash1", "field0")
		return r1
	`
	resHash1Field0Exists, err := rdb.Eval(ctx, luaScript_Hash1Field0Exist, []string{}).Result()
	if err != nil {
		log.Fatalf("redis hexists hash1 field0 error: %s", err)
	}
	log.Printf("redis hexists hash1 field0: %v", resHash1Field0Exists) // 0 or 1  //1代表有

	luaScript_Hash1 := `
		local r1 = ""
		r1 = redis.call("HGETALL", "hash1")
		return r1
	`
	resHash1, err := rdb.Eval(ctx, luaScript_Hash1, []string{}).Result()
	if err != nil {
		log.Fatalf("redis hgetall hash1 error: %s", err)
	}
	log.Printf("redis hgetall hash1: %v", resHash1) // [key1 value1 key2 value2 ...]
}
