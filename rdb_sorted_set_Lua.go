package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_sorted_set_ByLua(rdb *redis.Client, ctx context.Context) {
	log.Println()
	luaScript_Sort_Set1 := `
		local r1 = ""
		r1 = redis.call("ZRANGE", "sort_set1", 0, 10)
		return r1
	`
	resSort_Set1, err := rdb.Eval(ctx, luaScript_Sort_Set1, []string{}).Result()
	if err != nil {
		log.Fatalf("redis zrange sort_set1 error: %s", err)
	}
	log.Printf("redis zrange sort_set1: %v", resSort_Set1) //[value1 value2 ...]

	luaScript_Sort_Set1WithScores := `
		local r1 = ""
		r1 = redis.call("ZRANGE", "sort_set1", 0, 10, "withscores")
		return r1
	`
	resSort_Set1_WithScores, err := rdb.Eval(ctx, luaScript_Sort_Set1WithScores, []string{}).Result()
	if err != nil {
		log.Fatalf("redis zrangewithscores sort_set1 error: %s", err)
	}
	log.Printf("redis zrangewithscores sort_set1: %v", resSort_Set1_WithScores) //[value1 score1 value2 score2 ...]
}
