// Copyright 2018 liwenhua. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"sync"
	"time"
)

var lock sync.Mutex

type StringObj struct {
	Data map[string]*Context
}

func (c *StringObj) Set(key string, val interface{}) int {
	if c.Data == nil {
		c.Data = make(map[string]*Context)
	}

	data := &Context{val:val}
	c.Data[key] = data
	return REDIS_OK
}

func (c *StringObj) Get(key string) interface{} {
	if c.Data == nil {
		return nil
	}

	if data, ok := c.Data[key]; ok {
		if c.IsExpired(key, data.expireTime) {
			delete(c.Data, key)
			return nil
		}
		return c.Data[key].val
	} else {
		return nil
	}
}

func (c *StringObj) Incr(key string) int64 {
	return c.IncrBy(key, 1)
}

func (c *StringObj) IncrBy(key string, step int) int64 {
	lock.Lock()
	if c.Data == nil {
		if c.Data == nil {
			c.Data = make(map[string]*Context)
		}
		data := &Context{}
		data.val = step
		c.Data[key] = data
		c.Data[key].val = step
	} else {
		if _, ok := c.Data[key]; ok {
			c.Data[key].val = c.Data[key].val.(int64) + int64(step)
		} else {
			data := &Context{}
			data.val = step
			c.Data[key] = data
		}
	}
	lock.Unlock()
	return c.Data[key].val.(int64)
}

func (c *StringObj) Decr(key string) int64 {
	return c.DecrBy(key, 1)
}

func (c *StringObj) DecrBy(key string, step int) int64 {
	lock.Lock()
	if c.Data == nil {
		if c.Data == nil {
			c.Data = make(map[string]*Context)
		}
		data := &Context{}
		data.val = -step
		c.Data[key] = data
	} else {
		if _, ok := c.Data[key]; ok {
			c.Data[key].val = c.Data[key].val.(int64) - int64(step)
		} else {
			data := &Context{}
			data.val = -step
			c.Data[key] = data
		}

	}
	lock.Unlock()
	return c.Data[key].val.(int64)
}

func (c *StringObj) SetNx(key string, value interface{}) int {
	if c.Data == nil {
		c.Data = make(map[string]*Context)
		c.Data[key].val =  value
	}
	return REDIS_OK
}

func (c *StringObj) IsExpired(key string, curTime int64) bool {
	if c.Data[key].expireTime < curTime {
		return true
	}
	return false

}

func (c *StringObj) Expire(key string, expireTime int64) bool {
	if c.Data == nil {
		return false
	}

	if _,ok := c.Data[key]; ok {
		return false
	}
	curTime := time.Now().Unix()
	c.Data[key].expireTime = curTime + expireTime
	return true
}
