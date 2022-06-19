package	main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main()  {

	mux := new(http.ServeMux)

	mux.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[origin server] received request at: %s from IP : %s\n", time.Now(), r.RemoteAddr)
		fmt.Fprint(w, "origin server response, Method: ", r.Method)
		
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":8081"

	log.Println("running server at", server.Addr)
	log.Fatal(server.ListenAndServe())
}