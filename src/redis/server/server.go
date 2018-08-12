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
