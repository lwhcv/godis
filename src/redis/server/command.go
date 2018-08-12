package server

var StringObjCommand = map[string]int{
	"GET": 1, "SET": 2, "INCR": 0, "INCRBY": 1, "DECR": 0, "SETNX": 1, "SETEX": 1}

var HashObjCommand = map[string]int{"HGET": 2, "HSET": 1, "HDEL": 2, "HEXIST": 2, "HGETALL": 1, "HINCRBY": 3, "HKEYS": 1, "HLEN": 1, "HMGET": 3, "HMSET": 3, "HSETNX": 3}

var AllCommand = map[int]map[string]int{STRING: StringObjCommand, HASH: HashObjCommand}

func DoCommand(obj interface{}, oType int, command string, args []string) (interface{}, error) {
	index := 0
	// remove the first elements
	commandParams := append(args[:index], args[index+1:]...)
	cLen := len(commandParams)
	err := checkParam(command, oType, cLen)
	if err != nil {
		return nil, err
	}
	key := args[0]
	if oType == STRING {
		switch command {
		case "GET":
			return obj.(*StringObj).Get(key), nil
		case "SET":
			return obj.(*StringObj).Set(key, commandParams), nil
		case "INCR":
			return obj.(*StringObj).Incr(key), nil
		//case "INCRBY":
		//	return obj.(*StringObj).IncrBy(key, args[1].(int)), nil
		case "DECR":
			return obj.(*StringObj).Decr(key),nil
		case "SETNX":
			return obj.(*StringObj).SetNx(key, args[1]),nil
		default:
			return 0, nil

		}
	} else if oType == HASH {
		subKey := commandParams[0]
		switch command {
		case "HEGT":
			return obj.(*HashObj).HGet(key, subKey), nil
		case "HSET":
			return obj.(*HashObj).HSet(key, subKey, commandParams[1]), nil
		case "HDEL":
			return obj.(*HashObj).HDel(key, subKey), nil
		case "HEXIST":
			return obj.(*HashObj).HExists(key, subKey), nil
		case "HLEN":
			return obj.(*HashObj).HLen(key),nil
		case "HGETALL":
			return obj.(*HashObj).HGetAll(key),nil
		case "HSETNX":
			return obj.(*HashObj).HSetNx(key, subKey,commandParams[1]),nil
		default:
			return 0, nil

		}
	}

	return 0,nil
}

func checkParam(command string, cType int, cLen int) error {
	obj := AllCommand[cType]
	paramNum := obj[command]
	if cLen < paramNum {
		return CommandUseWrong
	}

	return nil
}
