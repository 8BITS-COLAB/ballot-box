package network

import (
	"context"
	"time"

	"github.com/perlin-network/noise"
)

func Ping(node *noise.Node, addr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	_, err := node.Ping(ctx, addr)

	if err != nil {
		return err
	}

	return nil
}
