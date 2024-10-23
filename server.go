package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "restaurant-service/restaurant-service/proto"
)

// Server implements the RestaurantServiceServer interface
type Server struct {
	pb.UnimplementedRestaurantServiceServer
	restaurants []*pb.Restaurant
}

// AddRestaurant adds a new restaurant to the list
func (s *Server) AddRestaurant(ctx context.Context, restaurant *pb.Restaurant) (*pb.Response, error) {
	s.restaurants = append(s.restaurants, restaurant)
	return &pb.Response{Success: true, Message: "Restaurant added successfully"}, nil
}

// GetAllRestaurants returns all restaurants
func (s *Server) GetAllRestaurants(ctx context.Context, empty *pb.Empty) (*pb.RestaurantList, error) {
	return &pb.RestaurantList{Restaurants: s.restaurants}, nil
}

// GetRestaurantByName finds a restaurant by name
func (s *Server) GetRestaurantByName(ctx context.Context, req *pb.NameRequest) (*pb.Restaurant, error) {
	for _, restaurant := range s.restaurants {
		if restaurant.Name == req.Name {
			return restaurant, nil
		}
	}
	return nil, fmt.Errorf("Restaurant not found")
}

// GetMenuItemsByCategory returns menu items by category
func (s *Server) GetMenuItemsByCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.MenuItemList, error) {
	var items []*pb.MenuItem
	for _, restaurant := range s.restaurants {
		for _, item := range restaurant.Menu {
			if item.Category == req.Category {
				items = append(items, item)
			}
		}
	}
	return &pb.MenuItemList{Items: items}, nil
}

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the RestaurantService server
	server := &Server{}
	pb.RegisterRestaurantServiceServer(grpcServer, server)

	// Start serving
	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
