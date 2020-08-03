package repo

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)
var ctx = context.Background()
// 声明一个全局的rdb变量
var RDB *redis.Client

// 初始化连接
func InitRedisClient() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = RDB.Ping(ctx).Result()
	return
}

func SetExp(key string,value interface{},duration time.Duration)error{
	return RDB.Set(ctx,key,value,duration).Err()
}

func Get(key string)(string,error){
	return RDB.Get(ctx,key).Result()
}