package smarthome

import (
	"context"
	"fmt"
	"time"

	"github.com/grandcat/zeroconf"
)

type resolver struct{}

func (r resolver) Init() error {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return err
	}
	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			fmt.Println(entry)
		}
		fmt.Println("No more entries")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err = resolver.Browse(ctx, "_services._udp", "local.", entries)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
