package std

import (
	"net/http"
	"testing"
)

// TestHttpServer 测试http服务端
func TestHttpServer(t *testing.T) {
	handler := func(respWriter http.ResponseWriter, req *http.Request) {
		// respWriter.Header().Set("Content-Type", "image/gif")
		respWriter.Header().Set("Content-Type", "text/plain")
		respWriter.Write([]byte("Hello Work!"))
	}
	mux := http.NewServeMux()
	mux.Handle("/lalala", http.HandlerFunc(handler))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
	select {}
}
