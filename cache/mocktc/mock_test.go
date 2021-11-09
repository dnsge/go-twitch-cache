package mocktc

import (
	"github.com/nicklaw5/helix"
	"testing"
)

func TestMockCache_GetUsers(t *testing.T) {
	cli := New([]helix.User{
		{
			ID:          "123",
			Login:       "user1",
			DisplayName: "User1",
		},
		{
			ID:          "456",
			Login:       "user2",
			DisplayName: "User2",
		},
		{
			ID:          "789",
			Login:       "user3",
			DisplayName: "User3",
		},
	}, nil, nil, nil)

	res, _ := cli.GetUsers(&helix.UsersParams{
		IDs:    []string{"123"},
		Logins: []string{"user3"},
	})

	user1 := res.ID("123")
	if user1 == nil {
		t.Fatalf("Got nil for user1")
	} else {
		if user1.DisplayName != "User1" {
			t.Fatalf("Got %q for user1 display name, expected \"User1\"", user1.DisplayName)
		}
	}

	user3 := res.ID("789")
	if user3 == nil {
		t.Fatalf("Got nil for user3")
	} else {
		if user3.DisplayName != "User3" {
			t.Fatalf("Got %q for user3 display name, expected \"User3\"", user3.DisplayName)
		}
	}
}

func TestMockCache_GetGames(t *testing.T) {
	cli := New(nil, []helix.Game{
		{
			ID:   "123",
			Name: "game1",
		},
		{
			ID:   "456",
			Name: "game2",
		},
		{
			ID:   "789",
			Name: "Cool game",
		},
	}, nil, nil)

	res, _ := cli.GetGames(&helix.GamesParams{
		IDs:   []string{"123"},
		Names: []string{"Cool game"},
	})

	game1 := res.ID("123")
	if game1 == nil {
		t.Fatalf("Got nil for game1")
	} else {
		if game1.Name != "game1" {
			t.Fatalf("Got %q for game1 name, expected \"game1\"", game1.Name)
		}
	}

	game3 := res.ID("789")
	if game3 == nil {
		t.Fatalf("Got nil for game3")
	} else {
		if game3.Name != "Cool game" {
			t.Fatalf("Got %q for game3 display name, expected \"Cool game\"", game3.Name)
		}
	}
}

func TestMockCache_GetStreams(t *testing.T) {
	cli := New(nil, nil, []helix.Stream{
		{
			UserID:   "123",
			UserName: "Streamer1",
		},
		{
			UserID:   "456",
			UserName: "Streamer2",
		},
		{
			UserID:   "789",
			UserName: "Streamer3",
		},
	}, nil)

	res, _ := cli.GetStreams(&helix.StreamsParams{
		UserIDs:    []string{"123"},
		UserLogins: []string{"streamer3"},
	})

	if len(res) != 2 {
		t.Fatalf("Expected 2 items in result, got %d", len(res))
	}
}

func TestMockCache_SearchChannels(t *testing.T) {
	cli := New(nil, nil, nil, []helix.Channel{
		{
			ID:          "123",
			DisplayName: "Channel1",
		},
		{
			ID:          "456",
			DisplayName: "Channel2",
		},
	})

	res, _ := cli.SearchChannels(&helix.SearchChannelsParams{
		Channel: "Channel2",
	})

	if len(res) != 1 {
		t.Fatalf("Expected 1 item in result, got %d", len(res))
	}

	if res[0].ID != "456" {
		t.Fatalf("Expected channel id 456, got %s", res[0].ID)
	}
}
