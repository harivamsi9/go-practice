/* KEY-VALUE DATABASE SERVER

$ curl -d 'hello' http://localhost:8080/k1
$ curl http://localhost:8080/k1
hello

$ curl -i http://localhost:8080/k2
404 not found

Limit value size to 1k
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type Server struct {
	db DB
}

type DB struct {
	m sync.Map
}

func (db *DB) Get(key string) []byte {
	val, ok := db.m.Load(key)
	if !ok {
		return nil
	}

	return val.([]byte)
}

func (db *DB) Set(key string, value []byte) {
	db.m.Store(key, value)
}

// POST /key Store request body as value
// GET /<key> Send back value, or 404 if key not found

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	// TODO
	switch r.Method {
	case http.MethodGet:
		s.getHandler(w, r)
		return

	case http.MethodPost:

		s.postHandler(w, r)
		return
	}

	http.Error(w, "bad method", http.StatusMethodNotAllowed)
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	value := s.db.Get(key)
	if value == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Write(value)
}

func (s *Server) postHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	defer r.Body.Close()
	rdr := io.LimitReader(r.Body, 1<<10)
	data, err := io.ReadAll(rdr)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	s.db.Set(key, data)
	fmt.Fprintf(w, "%s set\n", key)

}

func main() {
	// TODO: Routing & start server
	var s Server

	http.HandleFunc("/", s.Handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
