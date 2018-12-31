# Seahorse

Seahorse is a Redis client and Redis Queue manger.

## Installation

```
go get -u github.com/Kamva/seahorse
```

### Supported Commands

- [ ] String Types (key-value types)
    - [ ] APPEND
    - [ ] BITCOUNT
    - [ ] DECR
    - [ ] DECRBY
    - [x] GET
    - [ ] GETRANGE
    - [ ] GETSET
    - [ ] INCR
    - [ ] INCRBY
    - [ ] INCRBYFLOAT
    - [ ] MGET
    - [ ] MSET
    - [x] SET
    - [ ] SETRANGE
    - [ ] STRLEN
- [ ] Lists
    - [ ] BLPOP
    - [ ] BRPOP
    - [ ] BRPOPLPUSH
    - [ ] LINDEX
    - [ ] LLEN
    - [ ] LPOP
    - [x] LPUSH
    - [x] LRANGE
    - [ ] LREM
    - [ ] LSET
    - [ ] LTRIM
    - [ ] RPOP
    - [ ] RPOPLPUSH
    - [x] RPUSH
- [ ] Sets
    - [ ] SADD
    - [ ] SCARD
    - [ ] SDIFF
    - [ ] SDIFFSTORE
    - [ ] SINTER
    - [ ] SINTERSTORE
    - [ ] SISMEMBER
    - [ ] SMEMBERS
    - [ ] SMOVE
    - [ ] SPOP
    - [ ] SRANDMEMBER
    - [ ] SREM
    - [ ] SUNION
    - [ ] SUNIONSTORE
- [ ] Hashes
    - [ ] HDEL
    - [ ] HEXISTS
    - [ ] HGET
    - [ ] HGETALL
    - [ ] HINCRBY
    - [ ] HINCRBYFLOAT
    - [ ] HKEYS
    - [ ] HLEN
    - [ ] HMGET
    - [ ] HMSET
    - [ ] HSET
    - [ ] HSTRLEN
    - [ ] HVALS
- [ ] General
    - [ ] DEL
    - [ ] DUMP
    - [x] EXISTS
    - [ ] EXPIRE
    - [ ] EVAL
    - [ ] EXPIREAT
    - [x] FLUSHALL
    - [ ] KEYS
    - [ ] PERSIST
    - [ ] PEXPIRE
    - [ ] PEXPIREAT
    - [ ] PTTL
    - [ ] RANDOMKEY
    - [ ] RENAME
    - [ ] SCAN
    - [ ] SORT
    - [ ] TTL
    - [ ] TYPE
    - [ ] UNLINK
