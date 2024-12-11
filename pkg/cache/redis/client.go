package redis

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/8thgencore/microservice-common/pkg/logger"
	"github.com/8thgencore/microservice-common/pkg/logger/sl"
	"github.com/redis/go-redis/v9"
)

// ErrKeyNotFound is returned when a key is not found in a map or other data structure
var ErrKeyNotFound = errors.New("key not found")

type cacheClient struct {
	rdb *redis.Client
}

// NewClient creates client for Redis communication.
func NewClient(opt *redis.Options) *cacheClient {
	rdb := redis.NewClient(opt)
	return &cacheClient{rdb: rdb}
}

// String commands
func (c *cacheClient) Set(ctx context.Context, key string, value interface{}) error {
	if err := c.rdb.Set(ctx, key, value, 0).Err(); err != nil {
		logger.Error("unable to set key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

func (c *cacheClient) SetEx(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	if err := c.rdb.SetEx(ctx, key, value, duration).Err(); err != nil {
		logger.Error("unable to set key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

func (c *cacheClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return "", ErrKeyNotFound
		}
		logger.Error("unable to get key from the cache", slog.String("key", key))
		return "", err
	}

	return val, nil
}

func (c *cacheClient) Del(ctx context.Context, key string) error {
	if _, err := c.rdb.Del(ctx, key).Result(); err != nil {
		logger.Error("unable to del key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

func (c *cacheClient) DelAll(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil // No keys to delete
	}
	if _, err := c.rdb.Del(ctx, keys...).Result(); err != nil {
		logger.Error("unable to DelAll keys in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) Incr(ctx context.Context, key string) error {
	if err := c.rdb.Incr(ctx, key).Err(); err != nil {
		logger.Error("unable to incr key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

func (c *cacheClient) Decr(ctx context.Context, key string) error {
	if err := c.rdb.Decr(ctx, key).Err(); err != nil {
		logger.Error("unable to decr key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

func (c *cacheClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	expiresAt, err := c.rdb.TTL(ctx, key).Result()
	if err != nil {
		logger.Error("unable to ttl key in the cache", slog.String("key", key))
		return expiresAt, err
	}

	return expiresAt, nil
}

func (c *cacheClient) Expire(ctx context.Context, key string, duration time.Duration) error {
	if err := c.rdb.Expire(ctx, key, duration).Err(); err != nil {
		logger.Error("unable to expire key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

func (c *cacheClient) ExpireAt(ctx context.Context, key string, tm time.Time) error {
	if err := c.rdb.ExpireAt(ctx, key, tm).Err(); err != nil {
		logger.Error("unable to expire key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

// Hash commands
func (c *cacheClient) HSet(ctx context.Context, key, field string, value interface{}) error {
	if err := c.rdb.HSet(ctx, key, field, value).Err(); err != nil {
		logger.Error("unable to set field in the hash", slog.String("key", key), slog.String("field", field))
		return err
	}

	return nil
}

func (c *cacheClient) HGet(ctx context.Context, key, field string) (string, error) {
	result, err := c.rdb.HGet(ctx, key, field).Result()
	if err != nil {
		logger.Error("unable to get field from the hash", slog.String("key", key), slog.String("field", field))
		return "", err
	}

	return result, nil
}

func (c *cacheClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	result, err := c.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		logger.Error("unable to get all fields from the hash", slog.String("key", key))
		return nil, err
	}

	return result, nil
}

func (c *cacheClient) HIncrBy(ctx context.Context, key, field string, incr int64) error {
	if _, err := c.rdb.HIncrBy(ctx, key, field, incr).Result(); err != nil {
		logger.Error("unable to increment field in hash in the cache", slog.String("key", key), slog.String("field", field))
		return err
	}

	return nil
}

// List commands
func (c *cacheClient) LPush(ctx context.Context, key string, value interface{}) error {
	if err := c.rdb.LPush(ctx, key, value).Err(); err != nil {
		logger.Error("unable to lpush key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

func (c *cacheClient) LPushAll(ctx context.Context, key string, values ...interface{}) (int64, error) {
	val, err := c.rdb.LPush(ctx, key, values...).Result()
	if err != nil {
		logger.Error("unable to LPushAll key in the cache", slog.String("key", key))
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) LPop(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.LPop(ctx, key).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return "", ErrKeyNotFound
		}
		logger.Error("unable to lpop key in the cache", slog.String("key", key))
		return "", err
	}

	return val, nil
}

func (c *cacheClient) RPop(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.RPop(ctx, key).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return "", ErrKeyNotFound
		}
		logger.Error("unable to rpop key in the cache", slog.String("key", key))
		return "", err
	}

	return val, nil
}

func (c *cacheClient) LTrim(ctx context.Context, key string, start, stop int64) error {
	if err := c.rdb.LTrim(ctx, key, start, stop).Err(); err != nil {
		logger.Error("unable to ltrim key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

func (c *cacheClient) LLen(ctx context.Context, key string) (int64, error) {
	val, err := c.rdb.LLen(ctx, key).Result()
	if err != nil {
		logger.Error("unable to llen key in the cache", slog.String("key", key))
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) LRange(ctx context.Context, key string) ([]string, error) {
	val, err := c.rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		logger.Error("unable to lrange key in the cache", slog.String("key", key))
		return nil, err
	}

	return val, nil
}

// Set commands

func (c *cacheClient) SAdd(ctx context.Context, key string, value interface{}) (int64, error) {
	val, err := c.rdb.SAdd(ctx, key, value).Result()
	if err != nil {
		logger.Error("unable to sadd key in the cache", slog.String("key", key))
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SAddAll(ctx context.Context, key string, values ...interface{}) (int64, error) {
	val, err := c.rdb.SAdd(ctx, key, values...).Result()
	if err != nil {
		logger.Error("unable to saddAll key in the cache", slog.String("key", key))
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SRem(ctx context.Context, key string, value interface{}) (int64, error) {
	val, err := c.rdb.SRem(ctx, key, value).Result()
	if err != nil {
		logger.Error("unable to SRem key in the cache", slog.String("key", key))
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SCard(ctx context.Context, key string) (int64, error) {
	val, err := c.rdb.SCard(ctx, key).Result()
	if err != nil {
		logger.Error("unable to scard key in the cache", slog.String("key", key))
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SIsMember(ctx context.Context, key string, value interface{}) (bool, error) {
	val, err := c.rdb.SIsMember(ctx, key, value).Result()
	if err != nil {
		logger.Error("unable to SIsMember key in the cache", slog.String("key", key))
		return false, err
	}

	return val, nil
}

func (c *cacheClient) SMembers(ctx context.Context, key string) ([]string, error) {
	values, err := c.rdb.SMembers(ctx, key).Result()
	if err != nil {
		logger.Error("unable to SMembers key in the cache", slog.String("key", key))
		return nil, err
	}

	return values, nil
}

// Sorted Set commands
func (c *cacheClient) ZAdd(ctx context.Context, key string, value interface{}) error {
	if err := c.rdb.ZAdd(ctx, key, redis.Z{
		Score:  float64(time.Now().UnixMilli()),
		Member: value,
	}).Err(); err != nil {
		logger.Error("unable to zadd key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

func (c *cacheClient) ZAddWithScore(ctx context.Context, key string, score float64, value interface{}) error {
	if err := c.rdb.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: value,
	}).Err(); err != nil {
		logger.Error("unable to ZAddWithScore key in the cache", slog.String("key", key))
		return err
	}

	return nil
}

func (c *cacheClient) ZRem(ctx context.Context, key string, value interface{}) (int64, error) {
	val, err := c.rdb.ZRem(ctx, key, value).Result()
	if err != nil {
		logger.Error("unable to ZRem key in the cache", slog.String("key", key))
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) ZPopMin(ctx context.Context, key string, nb int64) ([]string, error) {
	val, err := c.rdb.ZPopMin(ctx, key, nb).Result()
	if err != nil {
		logger.Error("unable to zpopmin key in the cache", slog.String("key", key))
		return nil, err
	}
	var members []string
	for _, member := range val {
		members = append(members, member.Member.(string))
	}

	return members, nil
}

func (c *cacheClient) ZCount(ctx context.Context, key string) (int64, error) {
	val, err := c.rdb.ZCount(ctx, key, "-inf", "+inf").Result()
	if err != nil {
		logger.Error("unable to zcount key in the cache", slog.String("key", key))
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) ZRange(ctx context.Context, key string) ([]string, error) {
	val, err := c.rdb.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		logger.Error("unable to zrange key in the cache", slog.String("key", key))
		return nil, err
	}

	return val, nil
}

// Connection management
func (c *cacheClient) Ping(ctx context.Context) error {
	if err := c.rdb.Ping(ctx).Err(); err != nil {
		logger.Error("unable to ping redis", sl.Err(err))
		return err
	}

	return nil
}
