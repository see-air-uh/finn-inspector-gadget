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
	ID     string
	User   string
	Date   time.Time
	Module string
	Event  string
	Action string
	Data   string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		User:      payload.User,
		Date:      payload.Date,
		Module:    payload.Module,
		Event:     payload.Event,
		Action:    payload.Action,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("error writing to mongo.", err)
		return err
	}
	*resp = "Processed payload via RPC: " + payload.Action + " at: " + payload.Date.String()
	return nil
}
