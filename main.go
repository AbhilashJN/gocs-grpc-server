package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/AbhilashJN/gocs-grpc-server/services/heatmap"
	"github.com/AbhilashJN/gocs-grpc-server/services/map_details"
	"github.com/AbhilashJN/gocs-grpc-server/services/players_list"
	stats_summary "github.com/AbhilashJN/gocs-grpc-server/services/stats_summary"
	"github.com/joho/godotenv"
	grpc "google.golang.org/grpc"
)

type server struct {
	map_details.UnimplementedMapDetailsServiceServer
	players_list.UnimplementedPlayersListServiceServer
	stats_summary.UnimplementedStatsSummaryServiceServer
	heatmap.UnimplementedHeatmapServiceServer
}

func (*server) GetMapName(ctx context.Context, in *map_details.MapNameRequest) (*map_details.MapNameResponse, error) {
	demoPath := os.Getenv("DEMOPATH")
	return map_details.GetMapName(demoPath), nil
}

func (*server) GetPlayersList(ctx context.Context, in *players_list.PlayersListRequest) (*players_list.PlayersListResponse, error) {
	demoPath := os.Getenv("DEMOPATH")
	return players_list.PlayersListWrapper(demoPath), nil
}

func (s *server) GetStatsSummary(ctx context.Context, in *stats_summary.GetStatsSummaryRequest) (*stats_summary.GetStatsSummaryResponse, error) {
	demoPath := os.Getenv("DEMOPATH")
	player := in.GetPlayer()
	statType := in.GetStatType()
	response := stats_summary.StatsSummaryWrapper(demoPath, player, statType)
	return response, nil
}

func (*server) GetHeatmapPositions(ctx context.Context, in *heatmap.HeatmapRequest) (*heatmap.HeatmapResponse, error) {
	demoPath := os.Getenv("DEMOPATH")
	fmt.Println(demoPath, "asd")
	player := in.GetPlayer()
	return heatmap.PlayersListWrapper(demoPath, player), nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	stats_summary.RegisterStatsSummaryServiceServer(s, &server{})
	players_list.RegisterPlayersListServiceServer(s, &server{})
	heatmap.RegisterHeatmapServiceServer(s, &server{})
	map_details.RegisterMapDetailsServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
