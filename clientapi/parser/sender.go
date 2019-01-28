package parser

import (
	"context"
	"fmt"
	"time"

	t "github.com/nonemax/porto-transport"
)

// Sender is interface for sending port data
type Sender interface {
	SendPort(port []byte) error
}

// GRPCSender is type for sending port data with gRPC
type GRPCSender struct {
	c t.TransportClient
}

// NewSender return new JRPCSender
func NewSender(c t.TransportClient) GRPCSender {
	return GRPCSender{
		c: c,
	}
}

// SendPort send port data
func (s *GRPCSender) SendPort(bytePort []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rep, err := s.c.SendPort(ctx, &t.SendPortRequest{Portjson: bytePort})
	if err != nil {
		return err
	}
	if rep.Message != "Ok" {
		return fmt.Errorf("Wrong respomse from server")
	}
	return nil
}
