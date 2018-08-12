package server

import "time"

type HashObj struct {
	Data map[string]*HashData
}

type HashData struct {
	val map[string]interface{}
	expireTime int64
	len int
}

func (c *HashObj) HSet(key string, subKey string , val interface{}) int {
	if c.Data == nil {
		c.Data = make(map[string]*HashData)
	}
	if _,ok := c.Data[key]; ok {
		c.Data[key] = &HashData{}
	}

	c.Data[key].val[subKey] = val
	c.Data[key].len++
	return REDIS_OK
}

func (c *HashObj) HGet(key string, subKey string) interface{} {
	if c.Data == nil {
		return nil
	}
	data,ok := c.Data[key]
	if !ok {
		return nil
	}
	curTime := time.Now().Unix()
	if data.expireTime < curTime {
		delete(data.val, subKey)
		return nil
	}
	if _,ok := data.val[subKey]; !ok {
		return nil
	}
	return c.Data[key].val[subKey]
}

func (c *HashObj) HDel(key string, subKey string) int {
	if c.Data == nil {
		return REDIS_FAIL
	}
	data := c.Data[key]
	if data == nil || data.val == nil {
		return REDIS_FAIL
	} else {
		if data.val[subKey] == nil {
			return REDIS_FAIL
		}
		delete(c.Data[key].val, subKey)
		return REDIS_OK
	}

}

func (c *HashObj) HExists(key string, subKey string) bool {
	if c.Data == nil {
		return false
	}
	data := c.Data[key]
	if data == nil || data.val == nil {
		return false
	} else {
		if c.Data[key].val[subKey] == nil {
			return false
		}
		return true
	}
}

func (c *HashObj) HLen(key string) int {

	if c.Data == nil || c.Data[key] == nil {
		return REDIS_FAIL
	}

	return c.Data[key].len
}

func (c *HashObj) HGetAll(key string) interface{} {

	if c.Data == nil {
		return nil
	}
	data := c.Data[key]
	if data == nil {
		return nil
	}
	if data.val == nil {
		return nil
	}

	arrData := make([]interface{}, 0)
	for _,v := range data.val {
		arrData = append(arrData, v)
	}
	return arrData
}

func (c *HashObj) HSetNx(key string, subKey string, value interface{}) int {
	if c.Data == nil || c.Data[key] == nil {
		 return c.HSet(key, subKey, value)
	}

	if c.Data[key].val != nil {
		if c.Data[key].val[subKey] != nil {
			return REDIS_FAIL
		}
		c.Data[key].val[subKey] = value
		return REDIS_OK
	} else {
		c.Data[key].val = map[string]interface{}{subKey:value}
		return REDIS_OK
	}

}
