package storage

import "github.com/alicebob/miniredis"

type KeyValueStorage interface {
	SetAdd(k string, elems ...string) (int, error)
	SRem(k string, fields ...string) (int, error)
	Members(k string) ([]string, error)
	IsMember(k, v string) (bool, error)
	Set(k, v string) error
	Get(k string) (string, error)
	Close()
}

func NewMiniRedis() (KeyValueStorage, error) {
	return miniredis.Run()
}
