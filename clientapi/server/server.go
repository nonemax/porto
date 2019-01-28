package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	t "github.com/nonemax/porto-transport"
)

// HTTPServer describes http server
type HTTPServer struct {
	Addr string
	c    t.TransportClient
}

// New creats ne HTTPServer
func New(addr string, c t.TransportClient) (HTTPServer, error) {
	h := HTTPServer{
		Addr: addr,
		c:    c,
	}
	return h, nil
}

// Start is for starting new server
func (h *HTTPServer) Start() {
	router := httprouter.New()
	router.GET("/port/:unloc", h.GetPort)

	log.Fatal(http.ListenAndServe(h.Addr, router))
}

// GetPort is handler for get port request
func (h *HTTPServer) GetPort(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rep, err := h.c.GetPort(ctx, &t.GetPortRequest{Name: ps.ByName("unloc")})
	if err != nil {
		log.Println("Error in GetPort:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, User-Agent, Cache-Control, Authorization")
	w.Header().Add("Access-Control-Max-Age", "86400")
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(rep.Portjson)
}
