package tredis

import (
	"context"

	tserialize "github.com/octoposprime/op-be-shared/tool/serialize"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
}

func NewRedisClient(host string, port string, password string, db int) *RedisClient {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
	return &RedisClient{redisClient}
}

func (r *RedisClient) WriteHKey(ctx context.Context, hKey string, field string, data any) error {
	return r.HSet(ctx, hKey, field, tserialize.NewSerializer(data).ToJson()).Err()
}

func (r *RedisClient) ReadHKey(ctx context.Context, hKey string, field string, dataType any) (any, error) {
	hVal, err := r.HGet(ctx, hKey, field).Result()
	if err != nil {
		return nil, err
	}
	return tserialize.NewSerializer(dataType).FormJson(hVal), nil
}

func (r *RedisClient) DeleteHKey(ctx context.Context, hKey string, field string) error {
	return r.HDel(ctx, hKey, field).Err()
}
