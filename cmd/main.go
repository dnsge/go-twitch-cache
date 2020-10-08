package main

import (
	"flag"
	"github.com/dnsge/go-twitch-cache"
	"github.com/dnsge/go-twitch-cache/cache"
	"log"
	"net"
	"os"
	"time"
)

func getEnvStringOrFatal(key string) string {
	val, found := os.LookupEnv(key)
	if !found {
		log.Fatalf("Error: environment variable \"%s\" required\n", key)
	}
	return val
}

func main() {
	bindAddress := flag.String("address", ":50051", "Bind address")
	cacheExpiration := flag.Duration("expiration", time.Hour*24, "cache expiration time")
	cleanupInterval := flag.Duration("cleanup", time.Minute*30, "cache cleanup interval")
	flag.Parse()

	server := twitch.NewServer(&cache.ClientOptions{
		Credentials: cache.Credentials{
			ClientID:     getEnvStringOrFatal("TWITCH_CLIENT_ID"),
			ClientSecret: getEnvStringOrFatal("TWITCH_CLIENT_SECRET"),
		},
		Expiration:      *cacheExpiration,
		CleanupInterval: *cleanupInterval,
	})

	lis, err := net.Listen("tcp", *bindAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listening on %s\n", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
