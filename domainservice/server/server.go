package server

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/nonemax/porto-entity"
	t "github.com/nonemax/porto-transport"
	"github.com/nonemax/porto/domainservice/db"
	"google.golang.org/grpc"
)

// PortServer is a struct for server
type PortServer struct {
	Address string
	DB      db.DB
	s       *grpc.Server
}

// New creates nes PortServer
func New(addrs string, db db.DB) PortServer {
	return PortServer{
		Address: addrs,
		DB:      db,
	}
}

// Start is for starting server
func Start(server *PortServer) error {
	l, err := net.Listen("tcp", server.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	t.RegisterTransportServer(s, server)
	server.s = s
	return server.s.Serve(l)
}

// GetPort is for getting port from db
func (p *PortServer) GetPort(ctx context.Context, in *t.GetPortRequest) (*t.GetPortReply, error) {
	port, err := p.DB.GetPort(in.Name)
	if err != nil {
		log.Println("GetPort error:", err)
		return nil, err
	}
	bytePort, err := json.Marshal(port)
	if err != nil {
		log.Println("GetPort error:", err)
		return nil, err
	}
	return &t.GetPortReply{Portjson: bytePort}, nil
}

// SendPort is for saving port to db
func (p *PortServer) SendPort(ctx context.Context, in *t.SendPortRequest) (*t.SendPortReply, error) {
	newPort := entity.Port{}
	err := json.Unmarshal(in.Portjson, &newPort)
	if err != nil {
		log.Println("SendPort error:", err)
		return nil, err
	}
	err = p.DB.SavePort(newPort)
	if err != nil {
		log.Println("SendPort error:", err)
		return nil, err
	}
	return &t.SendPortReply{Message: "Ok"}, nil
}
