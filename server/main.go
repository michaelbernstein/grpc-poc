package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "panopticon/panopticon"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedPanopticonServer
}

// Mock data generator for Inference
func generateMockInference() *pb.Inference {
	return &pb.Inference{
		Timestamp:               timestamppb.Now(),
		BehaviorStartTimestamp:  timestamppb.Now(),
		BehaviorEndTimestamp:    timestamppb.Now(),
		Duration:                durationpb.New(time.Hour),
		IsBot:                   rand.Float32() < 0.5,
		ConfidenceScore:         rand.Float32(),
		BotCategory:             randomString(),
		RiskLevel:               randomString(),
		Reasoning:               []string{randomString()},
		ResponseAction:          randomString(),
		TraceId:                 randomString(),
		Geolocation:             randomGeolocation(),
		ConfidenceIntervalLower: rand.Float32(),
		ConfidenceIntervalUpper: rand.Float32(),
	}
}

func randomString() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, 10)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func randomGeolocation() string {
	return fmt.Sprintf("%f,%f", rand.Float64()*180-90, rand.Float64()*360-180)
}

func (s *server) Query(ctx context.Context, in *pb.BotQuery) (*pb.Inference, error) {
	return generateMockInference(), nil
}

func (s *server) ProcessBotQueries(stream pb.Panopticon_ProcessBotQueriesServer) error {
	for {
		_, err := stream.Recv()
		if err != nil {
			return err
		}
		if err := stream.Send(generateMockInference()); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPanopticonServer(s, &server{})
	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
