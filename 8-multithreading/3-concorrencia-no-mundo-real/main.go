package main

import (
	"fmt"
	"net/http"
	// "sync"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		// m.Unlock()
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("voce teve %d visitas ", number)))
	})
	http.ListenAndServe(":8080", nil)
}