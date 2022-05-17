package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const network, addr = "tcp4", "0.0.0.0:8000"

func main() {
	fmt.Println("Connection...")
	conn, err := net.Dial(network, addr)
	if err != nil {
		log.Print(err.Error())
		return
	}
	defer conn.Close()
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Your request (URL or keyword): ")
	for {
		query, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = conn.Write([]byte(query))
		if err != nil {
			fmt.Println(err)
			return
		}

		msg, err := io.ReadAll(conn)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(msg))
	}
}
