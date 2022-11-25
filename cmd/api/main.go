package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/see-air-uh/finn-inspector-gadget/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// _ "github.com/jackc/pgconn"
	// _ "github.com/jackc/pgx/v4"
	// _ "github.com/jackc/pgx/v4/stdlib"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongodb_container:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// connect to mongo
	mongoClient, err := connectToMongo()

	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// Register the RPC server
	// 			THIS rpc var is built in
	err = rpc.Register(new(RPCServer))
	go app.rpcListen()

	// start web server
	app.gRPCListen()

	// log.Printf("Starting service on port %s\n", webPort)
	// srv := &http.Server{
	// 	Addr:    fmt.Sprintf(":%s", webPort),
	// 	Handler: app.routes(),
	// }

	// err = srv.ListenAndServe()
	// if err != nil {
	// 	log.Panic()
	// }

}

func (app *Config) rpcListen() error {
	log.Println("Starting RPC server on port ", rpcPort)

	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}

// func (app *Config) serve() {
// 	srv := &http.Server{
// 		Addr:    fmt.Sprintf(":%s", webPort),
// 		Handler: app.routes(),
// 	}

// 	err := srv.ListenAndServe()
// 	if err != nil {
// 		log.Panic()
// 	}
// }

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "admin12345",
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}
	log.Println("Connected to mongo.")

	return c, nil
}
