package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rollout/rox-go/v5"
)

// Flags holds your feature flags
type Flags struct {
	EnableTutorial *rox.Flag
}

var flags = &Flags{
	EnableTutorial: rox.NewFlag(false),
}

func main() {
	// Initialize Rox
	options := rox.NewOptionsBuilder().Build()
	err := rox.Register("", flags)
	if err != nil {
		log.Fatalf("Failed to register flags: %v", err)
	}

	sdkKey := os.Getenv("ROX_SERVER_KEY")
	if sdkKey == "" {
		log.Fatal("ROX_SERVER_KEY environment variable is not set")
	}

	// Setup SDK
	rox.Setup(sdkKey, options)

	// Check the flag value
	if flags.EnableTutorial.IsEnabled(nil) {
		fmt.Println("✅ Feature flag is ENABLED!")
	} else {
		fmt.Println("🚫 Feature flag is DISABLED.")
	}
}
