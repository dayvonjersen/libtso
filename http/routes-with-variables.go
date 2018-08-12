package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request, vars map[string]string)

func exampleRoute(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	io.WriteString(w, "...")
}
func testRoute(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	io.WriteString(w, fmt.Sprintf("this %s a %s !!!", vars["is"], vars["test"]))
}
func desuRoute(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	io.WriteString(w, "desu~"+vars["wa"])
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
		7777,
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

	routes := map[string]handlerFunc{
		"/":                   exampleRoute,
		"/this/is/a/test":     exampleRoute,
		"/:this/:is/:a/:test": testRoute,
		"/this/:is/a/:test":   testRoute,
		"/desu/:wa":           desuRoute,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request verb:", r.Method)
		fmt.Println("request URL:", r.URL.Path)
		for t, fn := range routes {
			if pathMatch(r, t) {
				log.Println(req(r), "<- \033[32m200\033[0m OK")
				fn(w, r, pathVars(r, t))
				return
			}
		}
		notFoundHandler(w, r)
	})
	listenAddr := fmt.Sprintf("%s:%d", addr, port)
	if ssl {
		log.Println("listening on", listenAddr, "(HTTPS)")
		log.Fatalln(http.ListenAndServeTLS(listenAddr, certFile, keyFile, nil))
	} else {
		log.Println("listening on", listenAddr, "(HTTP)")
		log.Fatalln(http.ListenAndServe(listenAddr, nil))
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	log.Println(req(r), "<- \033[33m404\033[0m Not Found")
	io.WriteString(w, "404 v_v")
}

func req(r *http.Request) string {
	return fmt.Sprint(
		r.Host, " ", r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")], " ",
		// r.Header.Get("User-Agent"),
		"\n -> ", r.Method, " ", r.URL, "\n",
	)
}

func pathMatch(r *http.Request, schema string) bool {
	path := strings.Split(r.URL.Path, "/")

	for i, s := range strings.Split(schema, "/") {
		if len(path) <= i {
			return false
		}
		if len(s) > 0 && s[0] == ':' {
			continue
		}
		if path[i] != s {
			return false
		}
	}
	return true
}

func pathVars(r *http.Request, schema string) map[string]string {
	vars := map[string]string{}
	path := strings.Split(r.URL.Path, "/")

	for i, s := range strings.Split(schema, "/") {
		if len(s) > 0 && s[0] == ':' {
			k := s[1:]
			if i < len(path) {
				vars[k] = path[i]
			} else {
				vars[k] = ""
			}
		}
	}
	return vars
}
