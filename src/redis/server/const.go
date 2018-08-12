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
