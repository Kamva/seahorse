package seahorse

import "github.com/gomodule/redigo/redis"

// List redis wrapper for list operations
type List struct {
	Redis
	key string
}

// Index Returns the element at index index in the list stored at key. The index
// is zero-based, so 0 means the first element, 1 the second element and so on.
// Negative indices can be used to designate elements starting at the tail of the
// list. Here, -1 means the last element, -2 means the penultimate and so forth.
//
// When the value at key is not a list, an error is returned.
func (l List) Index(index int) (interface{}, error) {
	result, err := l.connection.Do("LINDEX", l.key, index)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Length Returns the length of the list stored at key. If key does not exist,
// it is interpreted as an empty list and 0 is returned. An error is returned
// when the value stored at key is not a list.
func (l List) Length() (int, error) {
	result, err := redis.Int(l.connection.Do("LLEN", l.key))
	if err != nil {
		return 0, err
	}

	return result, nil
}

// LeftPop Removes and returns the first element of the list stored at key.
func (l List) LeftPop() (interface{}, error) {
	result, err := l.connection.Do("LPOP", l.key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// LeftPush Insert all the specified values at the head of the list stored at
// key. If key does not exist, it is created as empty list before performing the
// push operations. When key holds a value that is not a list, an error is
// returned.
func (l List) LeftPush(value interface{}) (int, error) {
	result, err := redis.Int(l.connection.Do("LPUSH", l.key, value))
	if err != nil {
		return 0, err
	}

	return result, nil
}

// Range Returns the specified elements of the list stored at key. The offsets
// start and stop are zero-based indexes, with 0 being the first element of the
// list (the head of the list), 1 being the next element and so on.
func (l List) Range(start int, end int) ([]interface{}, error) {
	values, err := redis.Values(l.connection.Do("LRANGE", l.key, start, end))
	if err != nil {
		return nil, err
	}

	return values, nil
}

// Remove Removes the first count occurrences of elements equal to value from
// the list stored at key. The count argument influences the operation in the
// following ways:
//      count > 0: Remove elements equal to value moving from head to tail.
//      count < 0: Remove elements equal to value moving from tail to head.
//      count = 0: Remove all elements equal to value.
func (l List) Remove(value interface{}, occurrences int) (int, error) {
	result, err := redis.Int(l.connection.Do("LREM", l.key, occurrences, value))
	if err != nil {
		return 0, err
	}

	return result, nil
}

// Set Sets the list element at index to value.
func (l List) Set(index int, value interface{}) (bool, error) {
	result, err := l.connection.Do("LSET", l.key, index, value)
	if err != nil {
		return false, err
	}

	return result.(string) == "OK", nil
}

// RightPop Removes and returns the last element of the list stored at key.
func (l List) RightPop() (interface{}, error) {
	result, err := l.connection.Do("RPOP", l.key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// RightPopLeftPush Atomically returns and removes the last element (tail) of
// the list stored at source, and pushes the element at the first element (head)
// of the list stored at destination.
func (l List) RightPopLeftPush(newKey string) (interface{}, error) {
	result, err := l.connection.Do("RPOPLPUSH", l.key, newKey)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// RightPush Insert all the specified values at the tail of the list stored at
// key. If key does not exist, it is created as empty list before performing the
// push operation. When key holds a value that is not a list, an error is
// returned.
func (l List) RightPush(value interface{}) (int, error) {
	result, err := redis.Int(l.connection.Do("RPUSH", l.key, value))
	if err != nil {
		return 0, err
	}

	return result, nil
}
