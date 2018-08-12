package server

import(
	"redis/logs"
)

var log *logs.BLogger
func GetLogger() *logs.BLogger {
	return log
}

func InitLogger() {
	log = logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"/Users/liwenhua/log/applogs/godis.log"}`)
}
