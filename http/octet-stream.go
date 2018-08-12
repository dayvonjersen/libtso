package main

import "net/http"
import "time"
import "fmt"
import "log"

func main() {
	http.HandleFunc("/", octetstream)
	log.Fatal(http.ListenAndServe("0.0.0.0:6699", nil))
}

func octetstream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/octet-stream")
	for {
		<-time.After(time.Second * 2)
		fmt.Fprintf(w, "ping\n")
		if wf, ok := w.(http.Flusher); ok {
			wf.Flush()
		} else {
			log.Fatalln("no flsuh 4 u")
		}
	}
}
