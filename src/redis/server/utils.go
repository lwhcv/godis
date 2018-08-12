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
	"fmt"
	"strconv"
	"github.com/Unknwon/goconfig"
	"strings"
	"os"
	"errors"
)

func MustString(any interface{}) string {
	if any == nil {
		return ""
	}
	switch val := any.(type) {
	case int, uint, int64, uint64, uint32, int32, uint8, int8, int16, uint16:
		return fmt.Sprintf("%d", val)
	case string:
		return val
	case []byte:
		return string(val)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)

	default:
		return fmt.Sprintf("%v", val)
	}

	return ""
}

func InArray(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}

func FileExist(filename string) bool {
	if filename == "" {
		return false
	}
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}


func ReadMapFromIni(confpath string, args ...string) (arrMap map[string]string, err error) {
	if !FileExist(confpath) {
		arrMap = make(map[string]string)
		err = errors.New("file not exists. path:" + confpath)
		panic(err)
	}

	ini, err := goconfig.LoadConfigFile(confpath)
	if err != nil {
		panic(err)
	}

	keylist := ini.GetKeyList("")
	arrMap = make(map[string]string, len(keylist))

	keyprefix := ""
	if len(args) > 0 {
		keyprefix = strings.TrimSpace(args[0])
	}

	if keyprefix == "" {
		for _, key := range keylist {
			arrMap[key], _ = ini.GetValue("", key)
		}
	} else {
		for _, key := range keylist {
			arrMap[strings.TrimPrefix(key, keyprefix)], _ = ini.GetValue("", key)
		}
	}

	return
}
