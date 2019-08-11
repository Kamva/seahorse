package seahorse

import (
	"github.com/Kamva/shark"
	"github.com/gomodule/redigo/redis"
)

// Redis is a wrapper for redigo redis client
type Redis struct {
	connection redis.Conn
}

// Lists returns a wrapper to do all the operation of a list.
func (r Redis) Lists(key string) *List {
	return &List{redis: r, key: key}
}

// Sets returns a wrapper to do all the operation of a set.
func (r Redis) Sets(key string) *Set {
	return &Set{redis: r, key: key}
}

// Set creates a key-value and set the `value` for `key` and returns true if
// successfully created the key-value
func (r Redis) Set(key string, value string) (bool, error) {
	result, err := r.connection.Do("SET", key, value)
	if err != nil {
		return false, err
	}

	return result.(string) == "OK", nil
}

// Get returns the value that has been set for `key`
func (r Redis) Get(key string) (string, error) {
	value, err := redis.String(r.connection.Do("GET", key))
	if err != nil {
		return "", err
	}

	return value, nil
}

// Exists checks for `key` exists in the redis database
func (r Redis) Exists(key string) (bool, error) {
	result, err := redis.Int(r.connection.Do("EXISTS", key))
	if err != nil {
		return false, nil
	}

	return result > 0, nil
}

// FlushAll flushes all keys in the redis
func (r Redis) FlushAll() error {
	_, err := r.connection.Do("FLUSHALL")

	return err
}

// NewRedis initiate the redis object and returns it
func NewRedis(url string) *Redis {
	connection, err := redis.DialURL(url)
	shark.PanicIfError(err)

	return &Redis{
		connection: connection,
	}
}
