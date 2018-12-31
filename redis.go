package seahorse

import (
	"github.com/Kamva/shark"
	"github.com/gomodule/redigo/redis"
)

// Redis is a wrapper for redigo redis client
type Redis struct {
	connection redis.Conn
	issueCode  string
}

// Set creates a key-value and set the `value` for `key` and returns true if
// successfully created the key-value
func (r Redis) Set(key string, value string) bool {
	result, err := r.connection.Do("SET", key, value)
	shark.PanicIfError(err)

	return result.(string) == "OK"
}

// Get returns the value that has been set for `key`
func (r Redis) Get(key string) string {
	value, err := redis.String(r.connection.Do("GET", key))
	shark.PanicIfError(err)

	return value
}

// LPush push `value` into `key` list from left and returns length of list
func (r Redis) LPush(key string, value interface{}) int {
	result, err := redis.Int(r.connection.Do("LPUSH", key, value))
	shark.PanicIfError(err)

	return result
}

// RPush push `value` into `key` list from right and returns length of list
func (r Redis) RPush(key string, value interface{}) int {
	result, err := redis.Int(r.connection.Do("RPUSH", key, value))
	shark.PanicIfError(err)

	return result
}

// LRange gets the elements in `key` list from `start` index to `end` index
func (r Redis) LRange(key string, start int, end int) []interface{} {
	values, err := redis.Values(r.connection.Do("LRANGE", key, start, end))
	shark.PanicIfError(err)

	return values
}

// Exists checks for `key` exists in the redis database
func (r Redis) Exists(key string) bool {
	result, err := redis.Int(r.connection.Do("EXISTS", key))
	shark.PanicIfError(err)

	return result > 0
}

// FlushAll flushes all keys in the redis
func (r Redis) FlushAll() {
	_, err := r.connection.Do("FLUSHALL")
	shark.PanicIfError(err)
}

// NewRedis initiate the redis object and returns it
func NewRedis(url string, exceptionCode string) *Redis {
	connection, err := redis.DialURL(url)

	shark.PanicIfError(err)

	return &Redis{
		connection: connection,
		issueCode:  exceptionCode,
	}
}
