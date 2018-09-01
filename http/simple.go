package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

const tpl = `<!doctype html>
<html>
    <head>
        <meta charset='utf-8'>
        <title>{{.HelloWorld}}</title>
    </head>
    <body>
        <h1>{{.HelloWorld}}</h1>
    </body>
</html>`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-", r.Method, r.URL)
	setStatus := func(status int) {
		w.WriteHeader(status)
		log.Println("->", status, http.StatusText(status))
	}
	w.Header().Set("Content-Type", "text/html")

	data := &struct {
		HelloWorld string
	}{
		HelloWorld: "Hello, world.",
	}

	defer func() {
		t, err := template.New("").Parse(tpl)
		checkErr(err)
		buf := new(bytes.Buffer)
		checkErr(t.Execute(buf, data))
		io.Copy(w, buf)
	}()

	// do stuff here

	setStatus(200)

	return
}

func main() {
	var (
		addr              string
		port              int
		ssl               bool
		certFile, keyFile string
	)
	flag.StringVar(
		&addr,
		"addr",
		"",
		"leave blank for 0.0.0.0",
	)
	flag.IntVar(
		&port,
		"port",
		8080,
		"",
	)
	flag.BoolVar(
		&ssl,
		"ssl",
		false,
		"enable https",
	)
	flag.StringVar(
		&certFile,
		"cert",
		"cert.pem",
		"path to cert",
	)
	flag.StringVar(
		&keyFile,
		"key",
		"key.pem",
		"path to key",
	)
	flag.Parse()

	http.HandleFunc("/", indexHandler)

	listenAddr := fmt.Sprintf("%s:%d", addr, port)
	if ssl {
		log.Println("listening on", listenAddr, "(HTTPS) ...")
		log.Fatalln(http.ListenAndServeTLS(listenAddr, certFile, keyFile, nil))
	}
	log.Println("listening on", listenAddr, "(HTTP) ...")
	log.Fatalln(http.ListenAndServe(listenAddr, nil))
}

func checkErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
