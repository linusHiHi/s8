package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var Rdb *redis.Client

func RDInit(dbId int) *redis.Client {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "qwe12345",
		DB:       dbId,
	})
	return Rdb
}

func GetRedisValue(ctx context.Context, key string) (string, error) {
	GetKey := Rdb.Get(ctx, key)
	if GetKey.Err() != nil {
		return "", GetKey.Err()
	}
	return GetKey.Val(), nil
}

func SetRedisValue(ctx context.Context, key string, value string, expiration time.Duration) error {
	SetKV := Rdb.Set(ctx, key, value, expiration)
	return SetKV.Err()
}

type RedisSet struct {
	Id      int64
	Object  string
	Conn    *redis.Client
	Context context.Context
}

func NewRedisSet(context context.Context, Objet string, Id int64, Conn *redis.Client) *RedisSet {
	return &RedisSet{
		Id:      Id,
		Object:  Objet,
		Conn:    Conn,
		Context: context,
	}
}

func Set() {
	rs := NewRedisSet(context.Background(), "article:1", 0, Rdb)
	_, err := rs.Conn.SAdd(rs.Context, rs.Object, rs.Id).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var ctx = context.Background()
	client := RDInit(0)
	defer client.Close()
	err := SetRedisValue(ctx, "qwq", "a qwq", time.Hour)
	if err != nil {
		println()
	}
	value, _ := GetRedisValue(ctx, "qwq")
	print(value)
	Set()
}
