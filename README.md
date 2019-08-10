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
    - [x] LINDEX
    - [ ] LINSERT
    - [x] LLEN
    - [x] LPOP
    - [x] LPUSH
    - [x] LRANGE
    - [x] LREM
    - [x] LSET
    - [ ] LTRIM
    - [x] RPOP
    - [x] RPOPLPUSH
    - [x] RPUSH
- [ ] Sets
    - [x] SADD
    - [x] SCARD
    - [x] SDIFF
    - [ ] SDIFFSTORE
    - [ ] SINTER
    - [ ] SINTERSTORE
    - [x] SISMEMBER
    - [x] SMEMBERS
    - [x] SMOVE
    - [x] SPOP
    - [ ] SRANDMEMBER
    - [x] SREM
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
