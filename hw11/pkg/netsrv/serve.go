package netsrv

import (
	"bufio"
	"go2022/hw11/pkg/crawler"
	"go2022/hw11/pkg/index/hash"
	"log"
	"net"
)

func Handler(conn net.Conn, docs []crawler.Document) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}
		if len(msg) == 0 {
			conn.Write([]byte("Nothing found"))
			return
		}
		conn.Write([]byte("Results:\n"))
		index := hash.New()
		index.Add(docs)
		arr := index.Search(string(msg))
		log.Println("Found:")
		for i := range arr {
			for j := range docs {
				if arr[i] == docs[j].ID {
					log.Print(docs[j].URL)
					conn.Write([]byte(docs[j].URL))
				}
			}
		}
	}
}
