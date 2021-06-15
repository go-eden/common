package esync

import "sync/atomic"

type RMap struct {
	m atomic.Value
}

func (t *RMap) Get(key interface{}) (value interface{}) {
	m := t.m.Load()
	if m != nil {
		value = m.(map[interface{}]interface{})[key]
	}
	return nil
}

func (t *RMap) Set(key, value interface{}) {
	m := t.m.Load()
	if m != nil {
		value = m.(map[interface{}]interface{})[key]
	}
}

func (t *RMap) Del(key interface{}) {

}
