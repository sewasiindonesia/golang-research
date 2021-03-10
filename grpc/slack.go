package grpc

import (
	"context"
	slack "github.com/sewasiindonesia/golang-research/external"
	pb "github.com/sewasiindonesia/grpc/slack/proto"
	"log"
)

type Server struct {
}

func (s *Server) SendSlackMessage(ctx context.Context, req *pb.SendSlackMessageRequest) (*pb.SendSlackMessageResponse, error)  {
	sc := slack.SlackClient{
		WebHookUrl: "https://hooks.slack.com/services/blabla",
		Channel:    "product-tech",
	}
	sr := slack.SimpleSlackRequest{
		Text:      req.Message,
		IconEmoji: ":ghost:",
	}
	err := sc.SendSlackNotification(sr)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	return &pb.SendSlackMessageResponse{Success: true}, nil
}