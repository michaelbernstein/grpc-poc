package main

import (
	"context"
	"log"
	"time"

	pb "panopticon/panopticon"

	"google.golang.org/grpc"
)

func generateMockBotQuery() *pb.BotQuery {
	return &pb.BotQuery{
		Url:                    "http://example.com",
		UserAgent:              "Mozilla/5.0",
		Referrer:               "http://referrer.com",
		JA3Hash:                "example_hash",
		IpAddress:              "192.168.1.1",
		SessionId:              "example_session",
		HttpHeaders:            map[string]string{"Content-Type": "application/json"},
		RequestMethod:          "GET",
		RequestPayloadSize:     1234,
		RequestsCountInSession: 10,
		ReferralPath:           "/example",
		InteractionTime:        1.23,
	}
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPanopticonClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Single Query
	r, err := c.Query(ctx, generateMockBotQuery())
	if err != nil {
		log.Fatalf("could not query: %v", err)
	}
	log.Printf("Inference: %v", r)

	// Stream Queries
	stream, err := c.ProcessBotQueries(context.Background())
	if err != nil {
		log.Fatalf("could not stream: %v", err)
	}
	for i := 0; i < 5; i++ {
		queryBatch := &pb.BotQueryBatch{
			Queries: []*pb.BotQuery{generateMockBotQuery()},
		}
		if err := stream.Send(queryBatch); err != nil {
			log.Fatalf("could not send query batch: %v", err)
		}
		inference, err := stream.Recv()
		if err != nil {
			log.Fatalf("could not receive inference: %v", err)
		}
		log.Printf("Inference: %v", inference)
	}
	stream.CloseSend()
}
