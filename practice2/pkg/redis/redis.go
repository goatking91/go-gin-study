package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"net"
	"strings"

	"github.com/goatking91/go-gin-study/practice2/pkg/logger"
	"github.com/goatking91/go-gin-study/practice2/pkg/util"
)

var (
	ctx   = context.Background()
	Redis *redis.Client
)

func InitRedis() (ok bool) {
	ok = true
	env := &util.Env{EnvSource: &util.EnvGetter{}}
	logger.S.Info("Connecting Redis Server ...")

	address := net.JoinHostPort(env.GetString("REDIS_HOST"), env.GetString("REDIS_PORT"))
	logger.S.Infof("Connect Redis server. Addr:%s, Password:(***), DB:%d",
		address, env.GetInt("REDIS_DB"))

	// connect redis server
	Redis = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: env.GetString("REDIS_PASSWORD"),
		DB:       env.GetInt("REDIS_DB"),
	})

	if _, err := Ping(); err != nil {
		logger.S.Errorf("Fail connect to redis server. %v", err)
		ok = false
		return
	}

	//test redis
	testKeyValue()
	return
}

func Ping() (string, error) {
	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		logger.S.Errorf("redis ping fail. %v", err)
		return "Fail", err
	} else {
		logger.S.Debug("redis ping ok")
		return "Ok", nil
	}
}

// testKeyValue Redis 테스트. 테스트 순서 Set -> Get -> Delete
func testKeyValue() {

	logger.S.Info("Testing Redis Server ...")

	key := "Test key"
	value := "Test Key Value"

	logger.S.Debugf("Redis Test. Try Set key:(%s), value:(%s)", key, value)
	Set(key, value)

	logger.S.Debugf("Redis Test. Try Get key:(%s)", key)
	valueTmp, _ := Get(key)
	logger.S.Debugf("Redis Test. Result Get key:(%s), value:(%s)", key, valueTmp)

	logger.S.Debugf("Redis Test. Try Del key:(%s)", key)
	_ = Del(key)

}

// Set set redis key/value
func Set(key, value string) {
	err := Redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		logger.S.Errorf("redis Set error key:%s. %v", key, err)
	}
}

// Get get redis key
func Get(key string) (string, error) {
	val := ""
	val, err := Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		logger.S.Debugf("redis Get error no exist key:%s", key)
		err = nil
	} else if err != nil {
		logger.S.Errorf("redis Get error key:%s", key)
	}
	return val, err
}

// Del delete redis key
func Del(key string) error {
	_, err := Redis.Del(ctx, key).Result()
	if err == redis.Nil {
		logger.S.Debugf("redis Del error no exist key:%s", key)
		err = nil
	} else if err != nil {
		logger.S.Errorf("redis Del error key:%s", key)
	}
	return err
}

// DeleteFilter filter 로 검색해서 키 삭제.
//
// Parameters:
//  - `filter`: 검색 필터 (eg. prefix:*)
//
// Return:
//  - `int`: 삭제된 건수
func DeleteFilter(filter string) (int, error) {
	if strings.TrimSpace(filter) == "" {
		return 0, nil
	}

	cnt := 0

	// filter 로 scan
	iter := Redis.Scan(ctx, 0, filter, 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()
		if err2 := Del(key); err2 == nil {
			cnt++
		}
	}

	err := iter.Err()
	if err != nil {
		logger.S.Errorf("redis DeleteFilter. %v", err)
	}

	return cnt, err
}
