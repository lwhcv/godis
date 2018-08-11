package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "9999", "port")

type Resp struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println(">\nConnecting to " + *host + ":" + *port + "success")
	sendCommand(conn, "Set", "test", "Hello godis")
	ret := getReply(conn)
	fmt.Print(*host + ":" + *port + ">")
	fmt.Println(ret)
	sendCommand(conn, "Get", "test")
	ret = getReply(conn)
	fmt.Print(*host + ":" + *port + ">")
	fmt.Println(ret)
}

func sendCommand(conn net.Conn, args ...interface{}) error {

	msg := ""
	for _, v := range args {
		msg += v.(string) + "\\r\\n"
	}

	// marshal
	b, _ := json.Marshal(msg)
	writer := bufio.NewWriter(conn)
	_, err := writer.Write(b)
	if err != nil {
		fmt.Println("Error to send message because of ", err.Error())
		return err
	}
	// end with '\n' so that server can readline
	writer.Write([]byte("\n"))
	writer.Flush()
	return nil

}

func getReply(conn net.Conn) interface{} {
	reader := bufio.NewReader(conn)
	// read data
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Print("Error to read message because of ", err)
		return ""
	}
	// marshal
	var resp Resp
	json.Unmarshal(line, &resp)
	//fmt.Println("Status: ", resp.Status, " Content: ", resp.Data)
	return resp.Data

}
