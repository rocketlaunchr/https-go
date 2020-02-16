# https-go [![GoDoc](http://godoc.org/github.com/rocketlaunchr/https-go?status.svg)](http://godoc.org/github.com/rocketlaunchr/https-go)

Quickly create a self-signed Go HTTPS server.


## Example

```go
package main

import (
	"log"
	"net/http"
	"github.com/rocketlaunchr/https-go"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) })

	httpServer, _ := https.Server("8080", https.GenerateOptions{Host: "thecucumber.app"})
	log.Fatal(httpServer.ListenAndServeTLS("", ""))
}
```



### Extra Notes

Since the server is self-signed, some http clients will require extra configuration to allow self-signed certificates.


Other useful packages
------------

- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Statistics and data manipulation
- [dbq](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go
- [electron-alert](https://github.com/rocketlaunchr/electron-alert) - SweetAlert2 for Electron Applications
- [igo](https://github.com/rocketlaunchr/igo) - A Go transpiler with cool new syntax such as fordefer (defer for for-loops)
- [mysql-go](https://github.com/rocketlaunchr/mysql-go) - Properly cancel slow MySQL queries
- [react](https://github.com/rocketlaunchr/react) - Build front end applications using Go
- [remember-go](https://github.com/rocketlaunchr/remember-go) - Cache slow database queries
