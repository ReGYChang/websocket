package redis

import (
	"errors"
	"fmt"
	redisClient "github.com/gomodule/redigo/redis"
	"log"
	"time"
	"websocket/storage"
)

type redis struct{ pool *redisClient.Pool }

func New(host, port, password string) (storage.Service, error) {
	pool := &redisClient.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redisClient.Conn, error) {
			return redisClient.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		},
	}

	return &redis{pool}, nil
}

func (r *redis) Close() error {
	return r.pool.Close()
}

func (r *redis) Save(v interface{}) error {
	conn := r.pool.Get()
	defer func(conn redisClient.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	_, err := conn.Do("SET", "streams=btcusdt@aggTrade", v)
	if err != nil {
		log.Println("Redis Save failed: ", err)
		return err
	}

	return nil
}

func (r *redis) Load() (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	v, err := redisClient.String(conn.Do("GET", "streams=btcusdt@aggTrade"))
	if err != nil {
		return "", err
	} else if len(v) == 0 {
		return "", errors.New("data not found")
	}

	return v, nil
}
