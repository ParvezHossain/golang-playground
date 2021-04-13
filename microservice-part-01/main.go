package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// curl -v localhost:9090
	// curl -v -d "Parvez" localhost:9090: use -d to send data through URL
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		data, err := ioutil.ReadAll(r.Body)

		// If anny ERROR occurs the send error status-code
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Opps"))

			// Send custom error message
			// http.Error(rw, "Opps", http.StatusBadRequest)
			// set retun to terminate the operation, because HTTP ERROR does not terminate the operation, this will work with next REQUEST
			return
		}
		// log.Printf("Data: %s\n", data)
		// send response back to the user
		fmt.Fprintf(rw, "Hello, %s", data)
	})
	// curl -v localhost:9090/goodbye
	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World!")
	})
	// Creating the HTTP web server using DefaultServeMux
	http.ListenAndServe(":9090", nil)
}
