package server

import (
	"context"
	"fmt"
	"log"

	"github.com/Speshl/goremotecontrol_local/client/models"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	serialPort           *string
	baudRate             *int
	groundCommandChannel chan models.GroundControlData
}

func NewServer(serialPort *string, baudRate *int) *Server {
	return &Server{
		serialPort:           serialPort,
		baudRate:             baudRate,
		groundCommandChannel: make(chan models.GroundControlData, 10),
	}
}

func (s *Server) Start(ctx context.Context) error {
	log.Println("starting controller server...")
	defer log.Println("controller server stopped")

	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		return s.startSerial(ctx, s.groundCommandChannel)
	})

	err := errGroup.Wait()
	if err != nil {
		return fmt.Errorf("server error: %w", err)
	}
	return nil
}

func (s *Server) SendGroundCommand(command models.GroundControlData) {
	s.groundCommandChannel <- command
}
