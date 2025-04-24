package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_string_ByLua(rdb *redis.Client, ctx context.Context) {
	log.Println()
	luaScript_String1 := `
		local r1 = ""
		r1 = redis.call('GET', 'string1')
		return r1 
	`

	resString1, err := rdb.Eval(ctx, luaScript_String1, []string{}).Result()
	if err != nil {
		if err == redis.Nil {
			log.Println("redis get string1: nil")
		} else {
			log.Fatalf("redis get string1 error: %s", err)
		}
	} else {
		log.Printf("redis get string1: %s", resString1)
	}

	luaScript_String2 := `
		local r1 = ""
		r1 = redis.call('GET', 'string2')
		return r1
	`
	resString2, err := rdb.Eval(ctx, luaScript_String2, []string{}).Result()
	if err != nil {
		if err == redis.Nil {
			log.Println("redis get string2: nil")
		} else {
			log.Fatalf("redis get string2 error: %s", err)
		}
	} else {
		log.Fatalf("redis get string2: %s", resString2)
	}

}
