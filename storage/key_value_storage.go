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

var ErrKeyNotFound = miniredis.ErrKeyNotFound

func NewMiniRedis() (KeyValueStorage, error) {
	r, err := miniredis.Run()
	if err != nil {
		return nil, err
	}

	mrw := &miniRedisWrap{
		miniRedis: r,
	}

	return mrw, nil
}

type miniRedisWrap struct {
	miniRedis *miniredis.Miniredis
}

func (r *miniRedisWrap) SetAdd(k string, elems ...string) (int, error) {
	return r.miniRedis.SetAdd(k, elems...)
}

func (r *miniRedisWrap) SRem(k string, fields ...string) (int, error) {
	i, err := r.miniRedis.SRem(k, fields...)
	if err == miniredis.ErrKeyNotFound {
		return i, ErrKeyNotFound
	}

	return i, err
}

func (r *miniRedisWrap) Members(k string) ([]string, error) {
	return r.miniRedis.Members(k)
}

func (r *miniRedisWrap) IsMember(k, v string) (bool, error) {
	isMember, err := r.miniRedis.IsMember(k, v)
	if err == miniredis.ErrKeyNotFound {
		return isMember, ErrKeyNotFound
	}

	return isMember, err
}

func (r *miniRedisWrap) Set(k, v string) error {
	return r.miniRedis.Set(k, v)
}

func (r *miniRedisWrap) Get(k string) (string, error) {
	return r.miniRedis.Get(k)
}

func (r *miniRedisWrap) Close() {
	r.miniRedis.Close()
}
