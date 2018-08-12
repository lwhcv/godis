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
	"errors"
)

const STRING = 1
const LIST = 2
const HASH = 3

const REDIS_OK = 1 // success
const REDIS_FAIL = 0 // FAIL

var CommandInvalid = errors.New("command is invalid")
var CommandNotExist = errors.New("command not exist")
var CommandUseWrong = errors.New("command usage is not right")
