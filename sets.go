package seahorse

import "github.com/gomodule/redigo/redis"

// Set redis wrapper for set operations
type Set struct {
	redis Redis
	key   string
}

// Add the specified members to the set stored at key. Specified members that
// are already a member of this set are ignored. If key does not exist, a new
// set is created before adding the specified members.
func (s Set) Add(value interface{}) (bool, error) {
	result, err := redis.Bool(s.redis.connection.Do("SADD", s.key, value))
	if err != nil {
		return false, err
	}

	return result, nil
}

// Cardinality Returns the set cardinality (number of elements) of the set
// stored at key.
func (s Set) Cardinality() (int, error) {
	result, err := redis.Int(s.redis.connection.Do("SCARD", s.key))
	if err != nil {
		return 0, err
	}

	return result, nil
}

// Diff Returns the members of the set resulting from the difference between the
// first set and all the successive sets.
func (s Set) Diff(key2 string) ([]interface{}, error) {
	values, err := redis.Values(s.redis.connection.Do("SDIFF", s.key, key2))
	if err != nil {
		return nil, err
	}

	return values, nil
}

// IsMember Returns if member is a member of the set stored at key.
func (s Set) IsMember(value interface{}) (bool, error) {
	result, err := redis.Bool(s.redis.connection.Do("SISMEMBER", s.key, value))
	if err != nil {
		return false, err
	}

	return result, nil
}

// Members Returns all the members of the set value stored at key.
func (s Set) Members() ([]interface{}, error) {
	values, err := redis.Values(s.redis.connection.Do("SMEMBERS", s.key))
	if err != nil {
		return nil, err
	}

	return values, nil
}

// Move Move member from the set at source to the set at destination. This
// operation is atomic. In every given moment the element will appear to be a
// member of source or destination for other clients.
func (s Set) Move(otherKey string, value interface{}) (bool, error) {
	result, err := redis.Bool(s.redis.connection.Do("SMOVE", s.key, otherKey, value))
	if err != nil {
		return false, err
	}

	return result, nil
}

// Pop Removes and returns one random elements from the set value store at key.
func (s Set) Pop() (interface{}, error) {
	result, err := s.redis.connection.Do("SPOP", s.key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// PopMany Removes and returns specified random elements from the set value
// store at key.
func (s Set) PopMany(quantity int) ([]interface{}, error) {
	result, err := redis.Values(s.redis.connection.Do("SPOP", s.key, quantity))
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Remove Remove the specified members from the set stored at key. Specified
// members that are not a member of this set are ignored. If key does not exist,
// it is treated as an empty set and this command returns 0.
func (s Set) Remove(value interface{}) (int, error) {
	result, err := redis.Int(s.redis.connection.Do("SREM", s.key, value))
	if err != nil {
		return 0, err
	}

	return result, nil
}
