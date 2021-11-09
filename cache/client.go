package cache

import (
	"fmt"
	"github.com/nicklaw5/helix"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	userIDKeyPrefix   = "user:id:"
	userNameKeyPrefix = "user:name:"
	gameIDKeyPrefix   = "game:id:"
	gameNameKeyPrefix = "game:name:"
)

type Client interface {
	GetUsers(users *helix.UsersParams) (*UsersResult, error)
	GetGames(games *helix.GamesParams) (*GamesResult, error)
	GetUsersAndGames(users *helix.UsersParams, games *helix.GamesParams) (*UsersResult, *GamesResult, error)
	GetStreams(streams *helix.StreamsParams) ([]helix.Stream, error)
	SearchChannels(search *helix.SearchChannelsParams) ([]helix.Channel, error)
}

type ClientOptions struct {
	ClientID        string
	ClientSecret    string
	Expiration      time.Duration
	CleanupInterval time.Duration
}

type HelixCacheClient struct {
	options *ClientOptions

	client             *helix.Client
	tokenMutex         sync.RWMutex
	appToken           string
	appTokenExpiration time.Time

	cache *cache.Cache
}

func NewHelixCacheClient(options *ClientOptions) *HelixCacheClient {
	client, err := helix.NewClient(&helix.Options{
		ClientID:     options.ClientID,
		ClientSecret: options.ClientSecret,
		RedirectURI:  "",
	})

	if err != nil {
		panic(err)
	}

	return &HelixCacheClient{
		options:            options,
		client:             client,
		appToken:           "",
		appTokenExpiration: time.Now(),
		cache:              cache.New(options.Expiration, options.CleanupInterval),
	}
}

func (c *HelixCacheClient) CacheSize() uint32 {
	return uint32(c.cache.ItemCount())
}

func (c *HelixCacheClient) GetAppTokenClient() (*helix.Client, error) {
	var token string
	if c.tokenIsValid() {
		token = c.appToken
	} else {
		c.tokenMutex.Lock()
		err := c.generateAppToken()
		c.tokenMutex.Unlock()
		if err != nil {
			return nil, err
		}
		token = c.appToken
	}

	return helix.NewClient(&helix.Options{
		ClientID:       c.options.ClientID,
		ClientSecret:   c.options.ClientSecret,
		AppAccessToken: token,
		RedirectURI:    "",
	})
}

func (c *HelixCacheClient) tokenIsValid() bool {
	c.tokenMutex.RLock()
	defer c.tokenMutex.RUnlock()
	return c.appToken != "" && c.appTokenExpiration.After(time.Now())
}

func (c *HelixCacheClient) getAppTwitchClient() (*helix.Client, error) {
	if c.tokenIsValid() {
		return c.client, nil
	}

	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	err := c.generateAppToken()
	if err != nil {
		return nil, err
	}
	return c.client, nil
}

func (c *HelixCacheClient) setAppToken(token string, expiration time.Duration) {
	c.appToken = token
	c.appTokenExpiration = time.Now().Add(expiration)
	c.client.SetAppAccessToken(token)
}

func (c *HelixCacheClient) generateAppToken() error {
	// Generate new access token
	token, err := c.client.RequestAppAccessToken([]string{})
	if err != nil {
		return err
	}

	if token.StatusCode != http.StatusOK {
		log.Printf("Failed to generate app token: %v", token.ErrorMessage)
		return fmt.Errorf("failed to generate app token")
	}

	ttl := time.Duration(token.Data.ExpiresIn) * time.Second
	c.setAppToken(token.Data.AccessToken, ttl)
	return err
}

func (c *HelixCacheClient) requestNewAppToken() error {
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()
	return c.generateAppToken()
}

func (c *HelixCacheClient) cacheUsers(users []helix.User) {
	for _, user := range users {
		c.cache.Set(formatString(userIDKeyPrefix, user.ID), user, cache.DefaultExpiration)
		c.cache.Set(formatString(userNameKeyPrefix, user.Login), user, cache.DefaultExpiration)
	}
}

func (c *HelixCacheClient) getUsers(ids, names []string) ([]helix.User, error) {
	if len(ids)+len(names) > 100 {
		return nil, fmt.Errorf("100 maximum exceeded")
	}

	query := &helix.UsersParams{
		IDs:    ids,
		Logins: names,
	}

	client, err := c.getAppTwitchClient()
	if err != nil {
		return nil, err
	}

	helixRes, err := client.GetUsers(query)
	if err != nil {
		return nil, err
	}

	if helixRes.StatusCode == http.StatusUnauthorized {
		err := c.requestNewAppToken()
		if err != nil {
			return nil, err
		}
		helixRes, err = client.GetUsers(query)
		if err != nil {
			return nil, err
		}
	}

	if helixRes.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch users: %d %s", helixRes.StatusCode, helixRes.ErrorMessage)
		return nil, fmt.Errorf("failed to fetch users")
	}

	return helixRes.Data.Users, nil
}

func (c *HelixCacheClient) GetUsers(users *helix.UsersParams) (*UsersResult, error) {
	if len(users.IDs)+len(users.Logins) == 0 {
		return &UsersResult{[]helix.User{}}, nil
	}

	allIDs := removeDuplicates(users.IDs)
	allNames := removeDuplicates(users.Logins)

	var userResults []helix.User

	var requiredIDs []string
	for _, id := range allIDs {
		item, found := c.cache.Get(formatString(userIDKeyPrefix, id))
		if found {
			userResults = append(userResults, item.(helix.User))
		} else {
			requiredIDs = append(requiredIDs, id)
		}
	}

	var requiredNames []string
	for _, name := range allNames {
		if item, found := c.cache.Get(formatString(userNameKeyPrefix, name)); !found {
			requiredNames = append(requiredNames, name)
		} else {
			userResults = append(userResults, item.(helix.User))
		}
	}

	if len(requiredIDs)+len(requiredNames) > 0 {
		var needCache []helix.User
		for {
			a, b := calculateQuerySliceIndexes(requiredIDs, requiredNames)
			if a+b == 0 {
				break
			}

			res, err := c.getUsers(requiredIDs[:a], requiredNames[:b])
			if err != nil {
				return nil, err
			}
			needCache = append(needCache, res...)
			requiredIDs = requiredIDs[a:]
			requiredNames = requiredNames[b:]
		}
		userResults = append(userResults, needCache...)
		go c.cacheUsers(needCache)
	}

	return &UsersResult{userResults}, nil
}

func (c *HelixCacheClient) cacheGames(games []helix.Game) {
	for _, game := range games {
		c.cache.Set(formatString(gameIDKeyPrefix, game.ID), game, cache.DefaultExpiration)
		c.cache.Set(formatString(gameNameKeyPrefix, game.Name), game, cache.DefaultExpiration)
	}
}

func (c *HelixCacheClient) getGames(ids, names []string) ([]helix.Game, error) {
	if len(ids)+len(names) > 100 {
		return nil, fmt.Errorf("100 maximum exceeded")
	}

	query := &helix.GamesParams{
		IDs:   ids,
		Names: names,
	}

	client, err := c.getAppTwitchClient()
	if err != nil {
		return nil, err
	}

	helixRes, err := client.GetGames(query)
	if err != nil {
		return nil, err
	}

	if helixRes.StatusCode == http.StatusUnauthorized {
		err := c.requestNewAppToken()
		if err != nil {
			return nil, err
		}
		helixRes, err = client.GetGames(query)
		if err != nil {
			return nil, err
		}
	}

	if helixRes.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch games: %d %s", helixRes.StatusCode, helixRes.ErrorMessage)
		return nil, fmt.Errorf("failed to fetch games")
	}

	return helixRes.Data.Games, nil
}

func (c *HelixCacheClient) GetGames(games *helix.GamesParams) (*GamesResult, error) {
	if len(games.IDs)+len(games.Names) == 0 {
		return &GamesResult{[]helix.Game{}}, nil
	}

	allIDs := removeDuplicates(games.IDs)
	allNames := removeDuplicates(games.Names)

	var gameResults []helix.Game

	var requiredIDs []string
	for _, id := range allIDs {
		if item, found := c.cache.Get(formatString(gameIDKeyPrefix, id)); !found {
			requiredIDs = append(requiredIDs, id)
		} else {
			gameResults = append(gameResults, item.(helix.Game))
		}
	}

	var requiredNames []string
	for _, name := range allNames {
		if item, found := c.cache.Get(formatString(gameNameKeyPrefix, name)); !found {
			requiredNames = append(requiredNames, name)
		} else {
			gameResults = append(gameResults, item.(helix.Game))
		}
	}

	if len(requiredIDs)+len(requiredNames) > 0 {
		var needCache []helix.Game
		for {
			a, b := calculateQuerySliceIndexes(requiredIDs, requiredNames)
			if a+b == 0 {
				break
			}

			res, err := c.getGames(requiredIDs[:a], requiredNames[:b])
			if err != nil {
				return nil, err
			}
			needCache = append(needCache, res...)
			requiredIDs = requiredIDs[a:]
			requiredNames = requiredNames[b:]
		}
		gameResults = append(gameResults, needCache...)
		go c.cacheGames(needCache)
	}

	return &GamesResult{gameResults}, nil
}

// Utility function
func (c *HelixCacheClient) GetUsersAndGames(users *helix.UsersParams, games *helix.GamesParams) (*UsersResult, *GamesResult, error) {
	usersRes, err := c.GetUsers(users)
	if err != nil {
		return nil, nil, err
	}

	gamesRes, err := c.GetGames(games)
	if err != nil {
		return nil, nil, err
	}

	return usersRes, gamesRes, nil
}

func (c *HelixCacheClient) getStreams(ids, names []string) ([]helix.Stream, error) {
	if len(ids)+len(names) > 100 {
		return nil, fmt.Errorf("100 maximum exceeded")
	}

	query := &helix.StreamsParams{
		UserIDs:    ids,
		UserLogins: names,
	}

	client, err := c.getAppTwitchClient()
	if err != nil {
		return nil, err
	}

	helixRes, err := client.GetStreams(query)
	if err != nil {
		return nil, err
	}

	if helixRes.StatusCode == http.StatusUnauthorized {
		err := c.requestNewAppToken()
		if err != nil {
			return nil, err
		}
		helixRes, err = client.GetStreams(query)
		if err != nil {
			return nil, err
		}
	}

	if helixRes.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch streams: %d %s", helixRes.StatusCode, helixRes.ErrorMessage)
		return nil, fmt.Errorf("failed to fetch streams")
	}

	return helixRes.Data.Streams, nil
}

func (c *HelixCacheClient) GetStreams(streams *helix.StreamsParams) ([]helix.Stream, error) {
	if len(streams.UserIDs)+len(streams.UserLogins) == 0 {
		return []helix.Stream{}, nil
	}

	requiredIDs := removeDuplicates(streams.UserIDs)
	requiredNames := removeDuplicates(streams.UserLogins)

	var streamResults []helix.Stream

	if len(requiredIDs)+len(requiredNames) > 0 {
		for {
			a, b := calculateQuerySliceIndexes(requiredIDs, requiredNames)
			if a+b == 0 {
				break
			}

			res, err := c.getStreams(requiredIDs[:a], requiredNames[:b])
			if err != nil {
				return nil, err
			}
			streamResults = append(streamResults, res...)
			requiredIDs = requiredIDs[a:]
			requiredNames = requiredNames[b:]
		}
	}

	return streamResults, nil
}

func (c *HelixCacheClient) SearchChannels(search *helix.SearchChannelsParams) ([]helix.Channel, error) {
	client, err := c.getAppTwitchClient()
	if err != nil {
		return nil, err
	}

	helixRes, err := client.SearchChannels(search)
	if err != nil {
		return nil, err
	}

	if helixRes.StatusCode == http.StatusUnauthorized {
		err := c.requestNewAppToken()
		if err != nil {
			return nil, err
		}
		helixRes, err = client.SearchChannels(search)
		if err != nil {
			return nil, err
		}
	}

	if helixRes.StatusCode != http.StatusOK {
		log.Printf("Failed to search channels: %d %s", helixRes.StatusCode, helixRes.ErrorMessage)
		return nil, fmt.Errorf("failed to search channels")
	}

	return helixRes.Data.Channels, nil
}
