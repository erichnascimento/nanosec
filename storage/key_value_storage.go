package storage

import "github.com/alicebob/miniredis"

type KeyValueStorage interface {
	SetAdd(k string, elems ...interface{}) (int, error)
	SRem(k string, fields ...interface{}) (int, error)
	Members(k string) ([]interface{}, error)
	IsMember(k string, v interface{}) (bool, error)
	Set(k string, v interface{}) error
	Get(k string) (interface{}, error)
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

func (r *miniRedisWrap) SetAdd(k string, elems ...interface{}) (int, error) {
	return r.miniRedis.SetAdd(k, InterfaceListToStrList(elems)...)
}

func (r *miniRedisWrap) SRem(k string, fields ...interface{}) (int, error) {
	return r.miniRedis.SRem(k, InterfaceListToStrList(fields)...)
}

func (r *miniRedisWrap) Members(k string) ([]interface{}, error) {
	v, err := r.miniRedis.Members(k)
	return StrListToInterfaceList(v), err
}

func (r *miniRedisWrap) IsMember(k string, v interface{}) (bool, error) {
	return r.miniRedis.IsMember(k, InterfaceToStr(v))
}

func (r *miniRedisWrap) Set(k string, v interface{}) error {
	return r.miniRedis.Set(k, InterfaceToStr(v))
}

func (r *miniRedisWrap) Get(k string) (interface{}, error) {
	return r.miniRedis.Get(k)
}

func (r *miniRedisWrap) Close() {
	r.miniRedis.Close()
}

func InterfaceToStr(value interface{}) string {
	return value.(string)
}

func InterfaceListToStrList(list []interface{}) []string {
	l := make([]string, len(list))
	for i, v := range list {
		l[i] = InterfaceToStr(v)
	}

	return l
}

func StrToInterface(value string) (v interface{}) {
	v = value
	return
}

func StrListToInterfaceList(list []string) []interface{} {
	l := make([]interface{}, len(list))
	for i, v := range list {
		l[i] = StrToInterface(v)
	}

	return l
}
