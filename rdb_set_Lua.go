package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_set_ByLua(rdb *redis.Client, ctx context.Context) {
	log.Println()
	luaScript_Set1_Val0IsMember := `
		local r1 = ""
		r1 = redis.call("SISMEMBER", "set1", "set1_Val0")
		return r1
	`
	resSet1_Val0IsMember, err := rdb.Eval(ctx, luaScript_Set1_Val0IsMember, []string{}).Result()
	if err != nil {
		log.Fatalf("redis sismember set1 set1_Val0 error: %s", err)
	}
	log.Printf("redis sismember set1 set1_Val0: %v", resSet1_Val0IsMember) // 0 or 1  //1代表有

	luaScript_Set1 := `
		local r1 = ""
		r1 = redis.call("SMEMBERS", "set1")
		return r1
	`
	resSet1, err := rdb.Eval(ctx, luaScript_Set1, []string{}).Result()
	if err != nil {
		log.Fatalf("redis smembers set1 error: %s", err)
	}
	log.Printf("redis smembers set1: %v", resSet1) //[value1 value2 ...]
}
