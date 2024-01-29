package main

import (
	grpcClient "github.com/msharbaji/grpc-go-example/pkg/client"
	"context"
	"github.com/rs/zerolog/log"
)

func main() {
	client, _ := grpcClient.NewClient("localhost:50051", "1", "123456")
	users, _ := client.ListUsers(context.Background())
	log.Info().Msgf("%v", users)
}
