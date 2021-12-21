package redis

import (
	"github.com/go-redis/redis"
	"time"
)

var (
	client = &redisClient{}
)

type redisClient struct {
	c *redis.Client
}

func Test() string {
	return "asasd"
}

//Get redis client
func Initialize() *redisClient {
	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	client.c = c
	return client
}

//GetKey get key
func (client *redisClient) GetKey(key string) string {
	val, _ := client.c.Get(key).Result()
	return val
}

func (client *redisClient) GetAllKeys() []string {
	keys := client.c.Keys("*")
	return keys.Val()
}

//SetKey set key
func (client *redisClient) SetKey(key string, value string, expiration time.Duration) error {
	err := client.c.Set(key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

//Flush all keys
func (client *redisClient) FlushMemoryData() error {
	err := client.c.FlushAll().Err()
	if err != nil {
		return err
	}
	return nil
}
