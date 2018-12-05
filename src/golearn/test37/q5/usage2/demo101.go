package main

import (
	"net/http"
	"strings"
	"net/http/pprof"
	"log"
)

func main() {
	mux := http.NewServeMux()
	pathPrefx := "/d/pprof/"
	mux.HandleFunc(pathPrefx, func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(r.URL.Path, pathPrefx)
		if name != ""{
			pprof.Handler(name).ServeHTTP(w, r)
			return
		}
		pprof.Index(w, r)
	})

	mux.HandleFunc(pathPrefx+"cmdline", pprof.Cmdline)
	mux.HandleFunc(pathPrefx+"profile", pprof.Profile)
	mux.HandleFunc(pathPrefx+"symbox", pprof.Symbol)
	mux.HandleFunc(pathPrefx+"trace", pprof.Trace)

	server := http.Server{
		Addr: "localhost:8083",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err!= nil{
		if err == http.ErrServerClosed{
			log.Println("Http server closed.")
		}else{
			log.Printf("Http server error: %v\n", err)
		}
	}
}