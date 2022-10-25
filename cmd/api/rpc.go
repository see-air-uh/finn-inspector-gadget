package main

import (
	"context"
	"log"
	"time"

	"github.com/see-air-uh/finn-inspector-gadget/data"
)

// you need a specific type every time you set up an RPC server
type RPCServer struct{}

// define the payload you want to recieve from RPC
type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo.", err)
		return err
	}
	*resp = "Processed payload via RPC: " + payload.Name
	return nil
}
