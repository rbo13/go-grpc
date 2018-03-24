package main

import (
	"log"
	"net"

	"github.com/whaangbuu/grpc/server/blockchain"

	pb "github.com/whaangbuu/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Unable to listen on port 8080: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	srv.Serve(listener)
}

// Server implementation of the blockchain server.
type Server struct {
	Blockchain *blockchain.Blockchain
}

// AddBlock adds block to the chain
func (s *Server) AddBlock(ctx context.Context, in *pb.AddBlockRequest) (*pb.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data)

	return &pb.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

// GetBlockchain gets the blockchain
func (s *Server) GetBlockchain(ctx context.Context, in *pb.GetBlockChainRequest) (*pb.GetBlockchainResponse, error) {
	resp := new(pb.GetBlockchainResponse)

	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &pb.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash:          b.Hash,
			Data:          b.Data,
		})
	}
	return resp, nil
}
