# https-go [![GoDoc](http://godoc.org/github.com/rocketlaunchr/https-go?status.svg)](http://godoc.org/github.com/rocketlaunchr/https-go)

Quickly create a self-signed Go HTTPS server.


## Example

```go
package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) })

	httpServer, _ := Server("8080", GenerateOptions{Host: "thecucumber.app"})
	log.Fatal(httpServer.ListenAndServeTLS("", ""))
}
```



### Extra Notes

Since the server is self-signed, some http clients will require extra configuration to allow self-signed certificates.