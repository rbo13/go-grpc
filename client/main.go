package main

import (
	"flag"
	"log"
	"time"

	"golang.org/x/net/context"

	pb "github.com/whaangbuu/grpc/proto"
	"google.golang.org/grpc"
)

var client pb.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "Adds new block")
	listFlag := flag.Bool("list", false, "Lists blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial server: %v", err)
	}

	client = pb.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlag {
		getBlockchain()
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &pb.AddBlockRequest{
		Data: time.Now().String(),
	})

	if err != nil {
		log.Fatalf("Unable to add block: %v", err)
	}

	log.Printf("New Block Hash: %s\n", block.Hash)
}

func getBlockchain() {
	bc, err := client.GetBlockchain(context.Background(), &pb.GetBlockChainRequest{})

	if err != nil {
		log.Fatalf("Unable to get blockchain: %v", err)
	}

	log.Println("blocks:")
	for _, b := range bc.Blocks {
		log.Printf("Hash: %s, Previous Block Hash: %s, Data: %s", b.Hash, b.PrevBlockHash, b.Data)
	}
}
