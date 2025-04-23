package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_string(rdb *redis.Client, ctx context.Context) {
	log.Println()
	resString1, err := rdb.Get(ctx, "string1").Result()
	if err != nil {
		if err == redis.Nil {
			log.Println("redis get string1: nil")
		} else {
			log.Fatalf("redis get string1 error: %s", err)
		}
	} else {
		log.Printf("redis get string1: %s", resString1)
	}

	resString2, err := rdb.Get(ctx, "string2").Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("redis get string2: nil")
		} else {
			log.Fatalf("redis get string2 error: %s", err)
		}
	} else {
		log.Printf("redis get string2: %s", resString2)
	}

}
