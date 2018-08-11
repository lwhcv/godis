package server

var StringObjCommand = map[string]int{
	"GET":1, "SET":2, "INCR":0, "INCRBY":1, "DECR":0, "SETNX":1, "SETEX":1}

var HashObjCommand = map[string]int{"HGET":2, "HSET":1, "HGETALL":1}

var AllCommand = []interface{}{StringObjCommand, HashObjCommand}

func DoCommand(obj interface{}, cType int, command string, args []string) (interface{},error) {
	cLen := len(args)
	err := checkParam(command, cType, cLen)
	if err != nil {
		return nil,err
	}
	switch command {
	case "GET":
		return obj.(*StringObj).Get(args[0]),nil
	case "SET":
		return obj.(*StringObj).Set(args[0], args[1]),nil
	default:
		return 0,nil

	}

}


func checkParam(command string, cType int, cLen int) error {
	if cType == STRING {
		paramNum := StringObjCommand[command]
		if cLen < paramNum {
			return CommandUseWrong
		}
	} else if cType == HASH {
		paramNum := HashObjCommand[command]
		if cLen < paramNum {
			return CommandUseWrong
		}
	}

	return nil
}
