package main

import (
	"flag"
	"fmt"
	"github.com/dnsge/go-twitch-cache/cache"
	"github.com/dnsge/go-twitch-cache/rpc"
	"net"
	"os"
	"time"
)

func Fatalf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	os.Exit(1)
}

func getEnvStringOrFatal(key string) string {
	val, found := os.LookupEnv(key)
	if !found {
		Fatalf("Error: environment variable \"%s\" required\n", key)
	}
	return val
}

func main() {
	bindAddress := flag.String("address", ":50051", "Bind address")
	cacheExpiration := flag.Duration("expiration", time.Hour*24, "cache expiration time")
	cleanupInterval := flag.Duration("cleanup", time.Minute*30, "cache cleanup interval")
	flag.Parse()

	server := rpc.NewTwitchCacheServer(&cache.ClientOptions{
		ClientID:        getEnvStringOrFatal("TWITCH_CLIENT_ID"),
		ClientSecret:    getEnvStringOrFatal("TWITCH_CLIENT_SECRET"),
		Expiration:      *cacheExpiration,
		CleanupInterval: *cleanupInterval,
	})

	lis, err := net.Listen("tcp", *bindAddress)
	if err != nil {
		Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("listening on %s\n", lis.Addr())
	if err := server.Serve(lis); err != nil {
		Fatalf("failed to serve: %v", err)
	}
}
