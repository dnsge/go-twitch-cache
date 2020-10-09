package twitch

import (
	"context"
	"github.com/dnsge/go-twitch-cache/cache"
	pb "github.com/dnsge/go-twitch-cache/rpc"
	"github.com/nicklaw5/helix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime"
)

type server struct {
	pb.UnimplementedTwitchCacheServer
	client *cache.HelixCacheClient
}

func mapHelixUsersToPB(users *[]helix.User) *pb.Users {
	var all []*pb.User
	for _, user := range *users {
		all = append(all, &pb.User{
			ID:              user.ID,
			Login:           user.Login,
			DisplayName:     user.DisplayName,
			Description:     user.Description,
			ProfileImageURL: user.ProfileImageURL,
		})
	}
	return &pb.Users{Users: all}
}

func mapHelixGamesToPB(games *[]helix.Game) *pb.Games {
	var all []*pb.Game
	for _, game := range *games {
		all = append(all, &pb.Game{
			ID:        game.ID,
			Name:      game.Name,
			BoxArtURL: game.BoxArtURL,
		})
	}
	return &pb.Games{Games: all}
}

func mapHelixStreamsToPB(streams *[]helix.Stream) *pb.Streams {
	var all []*pb.Stream
	for _, stream := range *streams {
		all = append(all, &pb.Stream{
			ID:       stream.ID,
			UserID:   stream.UserID,
			UserName: stream.UserName,
			GameID:   stream.GameID,
			Title:    stream.Title,
		})
	}
	return &pb.Streams{Streams: all}
}

func (s *server) GetUsers(_ context.Context, query *pb.MultiQuery) (*pb.Users, error) {
	result, err := s.client.GetUsers(&helix.UsersParams{
		IDs:    query.IDs,
		Logins: query.Names,
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return mapHelixUsersToPB(result), nil
}

func (s *server) GetGames(_ context.Context, query *pb.MultiQuery) (*pb.Games, error) {
	result, err := s.client.GetGames(&helix.GamesParams{
		IDs:   query.IDs,
		Names: query.Names,
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return mapHelixGamesToPB(result), nil
}
func (s *server) GetStreams(_ context.Context, query *pb.MultiQuery) (*pb.Streams, error) {
	result, err := s.client.GetStreams(&helix.StreamsParams{
		UserIDs:    query.IDs,
		UserLogins: query.Names,
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return mapHelixStreamsToPB(result), nil
}

func (s *server) GetStatus(context.Context, *pb.StatusParams) (*pb.Status, error) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return &pb.Status{
		Alloc:     m.Alloc,
		Sys:       m.Sys,
		NumGC:     m.NumGC,
		CacheSize: s.client.CacheSize(),
	}, nil
}

func NewServer(options *cache.ClientOptions) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTwitchCacheServer(s, &server{
		client: cache.NewHelixCacheClient(options),
	})
	return s
}
