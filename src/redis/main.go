package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"redis/server"
	"strings"
)

var host = flag.String("host", "127.0.0.1", "host")
var port = flag.String("port", "9999", "port")

var stringObject = &server.StringObj{}
var hashObject = &server.HashObj{}

type Resp struct {
	Data   interface{} `json:"data"`
	Status int    `json:"status"`
}

func main() {
	flag.Parse()
	// listen
	l, err := net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on " + *host + ":" + *port)

	for {
		// accept client
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}

		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		go handleRequest(conn)
	}
}

// handler connection
func handleRequest(conn net.Conn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("Disconnected :" + ipStr)
		conn.Close()
	}()

	// init reader and writer
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		// readline end with "\n"
		b, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		// unmarshal
		var msg string
		json.Unmarshal(b, &msg)
		fmt.Println(msg)
		arrData := strings.Split(msg, "\\r\\n")
		fmt.Println(arrData)

		ret,err := ExecCommand(arrData)
		fmt.Println(stringObject)
		fmt.Println(ret)
		if err != nil {
			ret = err.Error()
		}

		// response Msg
		resp := Resp{
			Data:  ret,
			Status: 200,
		}

		r, _ := json.Marshal(resp)

		writer.Write(r)
		writer.Write([]byte("\n"))
		writer.Flush()
	}
}

func ExecCommand(args []string) (interface{}, error) {
	if len(args) < 1 {
		return nil, server.CommandInvalid
	}
	command := args[0]
	command = strings.ToUpper(args[0])
	objType := 0
	if _,ok := server.StringObjCommand[command]; ok {
		objType = server.STRING
	} else if _,ok := server.HashObjCommand[command]; ok {
		objType = server.HASH
	} else {
		return nil, server.CommandNotExist
	}

	index := 0
	// remove the first elements
	args = append(args[:index], args[index+1:]...)
	switch objType {
	case server.STRING:
		return server.DoCommand(stringObject, objType, command, args)
	case server.HASH:
		return server.DoCommand(hashObject, objType, command, args)
	}
	return 0, nil
}
