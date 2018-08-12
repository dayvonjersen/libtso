package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var i int = 0

func eventSource(w http.ResponseWriter, r *http.Request) {
	i++
	if i > 2 {
		i = 0
		return
	}
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	timeout := time.After(time.Second * 5)
	id := time.Now().Unix()
	log.Println("CONNECT", id)
	for {
		select {
		case <-timeout:
			log.Println("GOAWAY", id)
			return
		default:
			fmt.Fprintf(w, "id: %d\r\ndata: hello\r\n\r\n", id)
			<-time.After(time.Second)
		}
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		} else {
			log.Println("Can't flush")
		}
	}
}

func main() {
	http.HandleFunc("/", eventSource)
	log.Fatal(http.ListenAndServe(":777", nil))
}
