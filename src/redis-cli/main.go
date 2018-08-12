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
