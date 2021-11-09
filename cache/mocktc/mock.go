package mocktc

import (
	"github.com/dnsge/go-twitch-cache/cache"
	"github.com/nicklaw5/helix"
	"strings"
)

type mockCache struct {
	users    []helix.User
	games    []helix.Game
	streams  []helix.Stream
	channels []helix.Channel
}

func New(allUsers []helix.User, allGames []helix.Game, allStreams []helix.Stream, allChannels []helix.Channel) cache.Client {
	return &mockCache{
		users:    allUsers,
		games:    allGames,
		streams:  allStreams,
		channels: allChannels,
	}
}

func (m *mockCache) GetUsers(users *helix.UsersParams) (*cache.UsersResult, error) {
	var res []helix.User

	// Check IDs
	for _, id := range users.IDs {
		for i := range m.users {
			if m.users[i].ID == id {
				res = append(res, m.users[i])
				break
			}
		}
	}

	// Check logins
	for _, login := range users.Logins {
		login = strings.ToLower(login)
		for i := range m.users {
			if strings.ToLower(m.users[i].Login) == login {
				res = append(res, m.users[i])
				break
			}
		}
	}

	return &cache.UsersResult{Users: res}, nil
}

func (m *mockCache) GetGames(games *helix.GamesParams) (*cache.GamesResult, error) {
	var res []helix.Game

	// Check IDs
	for _, id := range games.IDs {
		for i := range m.games {
			if m.games[i].ID == id {
				res = append(res, m.games[i])
				break
			}
		}
	}

	// Check names
	for _, name := range games.Names {
		name = strings.ToLower(name)
		for i := range m.games {
			if strings.ToLower(m.games[i].Name) == name {
				res = append(res, m.games[i])
				break
			}
		}
	}

	return &cache.GamesResult{Games: res}, nil
}

func (m *mockCache) GetUsersAndGames(users *helix.UsersParams, games *helix.GamesParams) (*cache.UsersResult, *cache.GamesResult, error) {
	ur, _ := m.GetUsers(users)
	gr, _ := m.GetGames(games)
	return ur, gr, nil
}

// GetStreams only supports UserIDs and UserLogins fields
func (m *mockCache) GetStreams(streams *helix.StreamsParams) ([]helix.Stream, error) {
	var res []helix.Stream

	// Check IDs
	for _, id := range streams.UserIDs {
		for i := range m.streams {
			if m.streams[i].UserID == id {
				res = append(res, m.streams[i])
				break
			}
		}
	}

	// Check names
	for _, name := range streams.UserLogins {
		name = strings.ToLower(name)
		for i := range m.streams {
			if strings.ToLower(m.streams[i].UserName) == name {
				res = append(res, m.streams[i])
				break
			}
		}
	}

	return res, nil
}

func (m *mockCache) SearchChannels(search *helix.SearchChannelsParams) ([]helix.Channel, error) {
	name := strings.ToLower(search.Channel)
	for i := range m.channels {
		if name == strings.ToLower(m.channels[i].DisplayName) {
			return []helix.Channel{m.channels[i]}, nil
		}
	}
	return []helix.Channel{}, nil
}
