package rpc

import (
	"context"
	"github.com/dnsge/go-twitch-cache/cache"
	"github.com/nicklaw5/helix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime"
)

type server struct {
	UnimplementedTwitchCacheServer
	client *cache.HelixCacheClient
}

func mapHelixUsersToPB(users *[]helix.User) *Users {
	var all []*User
	for _, user := range *users {
		all = append(all, &User{
			ID:              user.ID,
			Login:           user.Login,
			DisplayName:     user.DisplayName,
			Description:     user.Description,
			ProfileImageURL: user.ProfileImageURL,
		})
	}
	return &Users{Users: all}
}

func mapHelixGamesToPB(games *[]helix.Game) *Games {
	var all []*Game
	for _, game := range *games {
		all = append(all, &Game{
			ID:        game.ID,
			Name:      game.Name,
			BoxArtURL: game.BoxArtURL,
		})
	}
	return &Games{Games: all}
}

func mapHelixStreamsToPB(streams *[]helix.Stream) *Streams {
	var all []*Stream
	for _, stream := range *streams {
		all = append(all, &Stream{
			ID:       stream.ID,
			UserID:   stream.UserID,
			UserName: stream.UserName,
			GameID:   stream.GameID,
			Title:    stream.Title,
		})
	}
	return &Streams{Streams: all}
}

func (s *server) GetUsersAndGames(_ context.Context, params *UsersAndGamesParams) (*UsersAndGames, error) {
	usersResult, err := s.client.GetUsers(&helix.UsersParams{
		IDs: params.UserIDs,
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	gamesResult, err := s.client.GetGames(&helix.GamesParams{
		IDs: params.GameIDs,
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &UsersAndGames{
		Users: mapHelixUsersToPB(usersResult),
		Games: mapHelixGamesToPB(gamesResult),
	}, nil
}

func (s *server) GetUsers(_ context.Context, query *MultiQuery) (*Users, error) {
	result, err := s.client.GetUsers(&helix.UsersParams{
		IDs:    query.IDs,
		Logins: query.Names,
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return mapHelixUsersToPB(result), nil
}

func (s *server) GetGames(_ context.Context, query *MultiQuery) (*Games, error) {
	result, err := s.client.GetGames(&helix.GamesParams{
		IDs:   query.IDs,
		Names: query.Names,
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return mapHelixGamesToPB(result), nil
}
func (s *server) GetStreams(_ context.Context, query *MultiQuery) (*Streams, error) {
	result, err := s.client.GetStreams(&helix.StreamsParams{
		UserIDs:    query.IDs,
		UserLogins: query.Names,
	})

	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return mapHelixStreamsToPB(result), nil
}

func (s *server) GetStatus(context.Context, *StatusParams) (*Status, error) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return &Status{
		Alloc:     m.Alloc,
		Sys:       m.Sys,
		NumGC:     m.NumGC,
		CacheSize: s.client.CacheSize(),
	}, nil
}

func NewTwitchCacheServer(options *cache.ClientOptions) *grpc.Server {
	s := grpc.NewServer()
	RegisterTwitchCacheServer(s, &server{
		client: cache.NewHelixCacheClient(options),
	})
	return s
}
