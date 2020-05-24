package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

type Counter struct {
	counter int
	mu sync.Mutex
}


func (s *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	hostname, err := os.Hostname()
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("i can't find myself in this world"))
	}
	resp := fmt.Sprintf("%s called %d times", hostname, s.counter)
	_, _ = w.Write([]byte(resp))
}

func main() {
	http.Handle("/", new(Counter))
	_ = http.ListenAndServe(":8080", nil)
}
