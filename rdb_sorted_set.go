package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func rdb_sorted_set(rdb *redis.Client, ctx context.Context) {
	log.Println()
	resSort_Set1, err := rdb.ZRange(ctx, "sort_set1", 0, 10).Result()
	if err != nil {
		log.Fatalf("redis zrange sort_set1 error: %s", err)
	}
	log.Printf("redis zrange sort_set1: %v", resSort_Set1) //[value1 value2 ...]

	resSort_Set1_WithScores, err := rdb.ZRangeWithScores(ctx, "sort_set1", 0, 10).Result()
	if err != nil {
		log.Fatalf("redis zrangewithscores sort_set1 error: %s", err)
	}
	log.Printf("redis zrangewithscore sort_set1: %v", resSort_Set1_WithScores) //[{score1 value1} {score2 value2} ...]
}
