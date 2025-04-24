package main

import (
	"context"
	"redis-get-test/configs"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func main() {
	redisHost := configs.EnvConfig.Redis.Host
	redisPort := configs.EnvConfig.Redis.Port
	redisDB := configs.EnvConfig.Redis.DB
	rdb, ctx := NewRedisClient(redisHost, redisPort, redisDB)
	defer rdb.Close()

	//string
	//rdb_string(rdb, ctx)
	rdb_string_ByLua(rdb, ctx)

	//hash
	//rdb_hash(rdb, ctx)
	rdb_hash_ByLua(rdb, ctx)

	//list
	//rdb_list(rdb, ctx)
	rdb_list_ByLua(rdb, ctx)

	//set
	//rdb_set(rdb, ctx)
	rdb_set_ByLua(rdb, ctx)

	//sorted set
	//rdb_sorted_set(rdb, ctx)
	rdb_sorted_set_ByLua(rdb, ctx)
}

func NewRedisClient(host string, port int, db int) (*redis.Client, context.Context) {
	var ctx = context.Background()
	opt := &redis.Options{
		Addr: host + ":" + strconv.Itoa(port),
		DB:   db,
	}
	rdb := redis.NewClient(opt)
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return rdb, ctx
}
