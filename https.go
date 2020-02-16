package https

import (
	"crypto/tls"
	"net/http"
)

// Server will create a self-signed https server.
//
// Example:
//
//  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) })
//
//  opts := GenerateOptions{Host: "domain.com"}
//  httpServer, _ := Server("8080", opts)
//  log.Fatal(httpServer.ListenAndServeTLS("", ""))
//
func Server(port string, opts GenerateOptions) (*http.Server, error) {

	pub, priv, err := GenerateKeys(opts)
	if err != nil {
		return nil, err
	}

	cert, err := tls.X509KeyPair(pub, priv)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:      ":" + port,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
	}, nil
}
