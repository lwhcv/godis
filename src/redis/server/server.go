package server

import "time"

type Context struct {
	val interface{}
	lastUpdateTime int64
	expireTime int64
}

func Cron(objects map[int]interface{}) {
	for {
		go ClearInvalidKeys(objects)
		go UpdateAof()
		time.Sleep(1)
	}
}

func ClearInvalidKeys(objects map[int]interface{}) {
	for oType,val := range objects {
		switch oType {
		case STRING:
			data := val.(*StringObj).Data
			for k,v := range data {
				curTime := time.Now().Unix()
				if v.expireTime <= curTime {
					delete(data, k)
				}
			}
		case LIST:
		case HASH:
			data := val.(*HashObj).Data
			for k,v := range data {
				curTime := time.Now().Unix()
				if v.expireTime <= curTime {
					delete(data, k)
				}
			}
		}
	}

}

func UpdateAof() {
	// use channel communicate
	// 根据配置文件来定时清理
}

func LoadDataFromFile() {

}
