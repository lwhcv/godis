package server

type HashObj struct {
	Data map[string]interface{}
}

// AddParam 添加请求参数
func (c *HashObj) Set(key string, val interface{}) int {
	if c.Data == nil {
		c.Data = make(map[string]interface{})
	}
	c.Data[key] = MustString(val)
	return REDIS_OK
}

func (c *HashObj) Get(key string) string {
	if c.Data == nil {
		return ""
	}
	return c.Data[key].(string)
}

func (c *HashObj) Incr(key string) int64 {
	lock.Lock()
	if c.Data == nil {
		c.Data[key] = 1
	} else {
		c.Data[key] = c.Data[key].(int64) + 1
	}
	lock.Unlock()
	return c.Data[key].(int64)
}

func (c *HashObj) IncrBy(key string, step int) int64 {
	lock.Lock()
	if c.Data == nil {
		c.Data[key] = step
	} else {
		c.Data[key] = c.Data[key].(int64) + int64(step)
	}
	lock.Unlock()
	return c.Data[key].(int64)
}

func (c *HashObj) Decr(key string) int64 {
	lock.Lock()
	if c.Data == nil {
		c.Data[key] = -1
	} else {
		c.Data[key] = c.Data[key].(int64) - 1
	}
	lock.Unlock()
	return c.Data[key].(int64)
}

func (c *HashObj) DecrBy(key string, step int) int64 {
	lock.Lock()
	if c.Data == nil {
		c.Data[key] = -step
	} else {
		c.Data[key] = c.Data[key].(int64) - int64(step)
	}
	lock.Unlock()
	return c.Data[key].(int64)
}

func (c *HashObj) SetNx(key string, value interface{}) {
	if c.Data == nil {
		c.Data = make(map[string]interface{})
		c.Data[key] = value
	}
}
