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
	ethereumUrl    = flag.String("url", "", "Ethereum URL to connect to")
	deployContract = flag.Bool("deploy", false, "Deploys the Store contract")
	passphrase     = flag.String("passphrase", "", "Passphrase of the minter's private key")
	pathToKey      = flag.String("key-path", "", "Path to the minter's private key")
	port           = flag.String("port", ":3005", "HTTP Port to listen on")
)

const (
	contractAddress      = ""
	goerliTestnetChainID = 5
)

type Service struct {
	client *ethclient.Client
	store  *bindings.Store
}

func main() {
	flag.Parse()

	ctx := context.Background()

	client, err := ethclient.Dial(*ethereumUrl)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	fmt.Println("Connected to Ethereum")

	service := &Service{client: client, store: nil}

	if *deployContract {
		err := service.deploy(ctx, *pathToKey, *passphrase)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	// Start the server if we aren't deploying
	store, err := bindings.NewStore(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Failed to initialize Store contract: %v", err)
	}
	service.store = store
	fmt.Println("💖 Loaded contract")

	router := mux.NewRouter()
	router.Use(PanicRecoveryMiddleware)
	router.HandleFunc("/set", http.HandlerFunc(service.handleSetItem)).Methods("POST")
	router.HandleFunc("/{key}", http.HandlerFunc(service.handleGetItem)).Methods("GET")

	fmt.Printf(":) Running dante on %v\n", *port)

	log.Fatal(http.ListenAndServe(*port, router))
}
