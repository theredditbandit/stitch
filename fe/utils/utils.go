package utils

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"os"
	"strings"
)

func IsFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return true, err
}

func IsFileSupported(path string) (bool, error) {
	supportedFiletypes := map[string]bool{"txt": true, "pdf": true}
	splitPath := strings.Split(path, ".")
	extension := strings.ToLower(splitPath[1])
	support, ok := supportedFiletypes[extension]
	if !ok {
		return false, errors.New("File type not supported")
	}
	return support, nil
}

func ConnectToRedis(ctx context.Context, cancel context.CancelFunc, rdbChan chan *redis.Client) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		rdbChan <- nil
	} else {
		rdbChan <- rdb
		cancel()
	}
}
