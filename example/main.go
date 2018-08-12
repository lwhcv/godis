package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"encoding/json"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "9999", "port")

type Msg struct {
	Data string `json:"data"`
	Type int    `json:"type"`
}

type Resp struct {
	Data string `json:"data"`
	Status int  `json:"status"`
}

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)
	// read and write
	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	// write 10 data
	for i := 10; i > 0; i-- {
		d := "hello " + strconv.Itoa(i)
		msg := Msg{
			Data: d,
			Type: 1,
		}
		// marsha
		b, _ := json.Marshal(msg)
		writer := bufio.NewWriter(conn)
		_, e := writer.Write(b)
		//_, e := conn.Write(b)
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}

		writer.Write([]byte("\n"))
		writer.Flush()
	}
	fmt.Println("Write Done!")
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	// read data
	for i := 1; i <= 10; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			return
		}
		// unmarshal
		var resp Resp
		json.Unmarshal(line, &resp)
		fmt.Println("Status: ", resp.Status, " Content: ", resp.Data)
	}
	fmt.Println("Read Done!")
}
