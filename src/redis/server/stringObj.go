package server

import (
	"sync"
)

var lock sync.Mutex

type StringObj struct {
	Data map[string]interface{}
}

func (c *StringObj) Set(key string, val interface{}) int {
	if c.Data == nil {
		c.Data = make(map[string]interface{})
	}
	c.Data[key] = MustString(val)
	return REDIS_OK
}

func (c *StringObj) Get(key string) string {
	if c.Data == nil {
		return ""
	}
	return c.Data[key].(string)
}

func (c *StringObj) Incr(key string) int64 {
	lock.Lock()
	if c.Data == nil {
		c.Data[key] = 1
	} else {
		c.Data[key] = c.Data[key].(int64) + 1
	}
	lock.Unlock()
	return c.Data[key].(int64)
}

func (c *StringObj) IncrBy(key string, step int) int64 {
	lock.Lock()
	if c.Data == nil {
		c.Data[key] = step
	} else {
		c.Data[key] = c.Data[key].(int64) + int64(step)
	}
	lock.Unlock()
	return c.Data[key].(int64)
}

func (c *StringObj) Decr(key string) int64 {
	lock.Lock()
	if c.Data == nil {
		c.Data[key] = -1
	} else {
		c.Data[key] = c.Data[key].(int64) - 1
	}
	lock.Unlock()
	return c.Data[key].(int64)
}

func (c *StringObj) DecrBy(key string, step int) int64 {
	lock.Lock()
	if c.Data == nil {
		c.Data[key] = -step
	} else {
		c.Data[key] = c.Data[key].(int64) - int64(step)
	}
	lock.Unlock()
	return c.Data[key].(int64)
}

func (c *StringObj) SetNx(key string, value interface{}) {
	if c.Data == nil {
		c.Data = make(map[string]interface{})
		c.Data[key] = value
	}
}
