package cache

import (
	"github.com/nicklaw5/helix"
	"strings"
)

type UsersResult struct {
	Users []helix.User
}

func (u *UsersResult) ID(id string) *helix.User {
	for _, user := range u.Users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func (u *UsersResult) Name(name string) *helix.User {
	name = strings.ToLower(name)
	for _, user := range u.Users {
		if strings.ToLower(user.Login) == name {
			return &user
		}
	}
	return nil
}

type GamesResult struct {
	Games []helix.Game
}

func (g *GamesResult) ID(id string) *helix.Game {
	for _, game := range g.Games {
		if game.ID == id {
			return &game
		}
	}
	return nil
}

func (g *GamesResult) Name(name string) *helix.Game {
	name = strings.ToLower(name)
	for _, game := range g.Games {
		if strings.ToLower(game.Name) == name {
			return &game
		}
	}
	return nil
}

