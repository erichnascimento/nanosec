package storage

import "github.com/alicebob/miniredis"

type KeyValueStorage interface {
	SetAdd(k string, elems ...string) (int, error)
	SRem(k string, fields ...string) (int, error)
	IsMember(k, v string) (bool, error)
	Close()
}

func NewMiniRedis() (KeyValueStorage, error) {
	return miniredis.Run()
}
