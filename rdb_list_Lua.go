package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_list_ByLua(rdb *redis.Client, ctx context.Context) {
	log.Println()
	luaScript_List1 := `
		local r1 = ""
		r1 = redis.call("LRANGE", "list1", 0, 2)
		return r1
	`
	resList1, err := rdb.Eval(ctx, luaScript_List1, []string{}).Result()
	if err != nil {
		log.Fatalf("redis lrange list1 error: %s", err)
	}
	log.Printf("redis lrange list1: %s", resList1) //[value1 value2 ...]
}
