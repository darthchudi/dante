package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/darthchudi/dante/bindings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	deployContract = flag.Bool("deploy", false, "Deploy store contract")
	ethereumUrl = flag.String("url", "", "Ethereum URL to connect to")
	port = flag.String("port", ":3005", "HTTP Port to listen on")
)

const (
	contractAddress = ""
)

type Service struct {
	client *ethclient.Client
	store *bindings.Store
}


func main () {
	flag.Parse()

	ctx := context.Background()

	client, err := ethclient.Dial(*ethereumUrl)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	fmt.Println("Connected to Ethereum")

	store, err := bindings.NewStore(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Failed to initialize Store contract: %v", err)
	}

	service := &Service{client: client, store: store}

	if *deployContract {
		service.deploy(ctx)
		fmt.Println("💖 Deployed contract")
		os.Exit(0)
	}

	router := mux.NewRouter()
	router.Use(PanicRecoveryMiddleware)
	router.HandleFunc("/set", http.HandlerFunc(service.handleSetItem)).Methods("POST")
	router.HandleFunc("/{key}", http.HandlerFunc(service.handleGetItem)).Methods("GET")

	fmt.Printf(":) Running dante on %v\n", *port)

	log.Fatal(http.ListenAndServe(*port, router))
}