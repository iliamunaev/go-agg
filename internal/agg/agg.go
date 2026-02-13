package agg

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Aggregator struct {
	Client            *http.Client
	MaxInflight       int
	PerRequestTimeout time.Duration
}

func New(client *http.Client, maxInflight int, perRequestTimeout time.Duration) *Aggregator {
	if client == nil {
		client = http.DefaultClient
	}
	if maxInflight <= 0 {
		maxInflight = 8
	}
	if perRequestTimeout <= 0 {
		perRequestTimeout = 2 * time.Second
	}
	return &Aggregator{
		Client:            client,
		MaxInflight:       maxInflight,
		PerRequestTimeout: perRequestTimeout,
	}
}

// Placeholder: temp
func (a *Aggregator) Aggregate(ctx context.Context, urls []string) error {
	if len(urls) == 0 {
		return fmt.Errorf("no urls")
	}
	return nil
}