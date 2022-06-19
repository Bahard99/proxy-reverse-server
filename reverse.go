package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main()  {

	// define origin server URl
	originServerURL, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	reverseProxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[reverse proxy server] received request at: %s from IP : %s\n", time.Now(), r.RemoteAddr)

		// set req Host, URL and Requeust URI to forward a request to the origin server
		r.Host 			= originServerURL.Host
		r.URL.Host 		= originServerURL.Host
		r.URL.Scheme 	= originServerURL.Scheme
		r.RequestURI	= ""

		// save the response from origin server
		originServerRes, err := http.DefaultClient.Do(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// return response to the client
		w.WriteHeader(http.StatusOK)
		io.Copy(w, originServerRes.Body)
	})

	log.Fatal(http.ListenAndServe(":8080", reverseProxy))
}
