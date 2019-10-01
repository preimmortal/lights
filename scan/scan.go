package scan

import (
	"context"
	"time"

	"github.com/Ullaakut/nmap"
	"github.com/preimmortal/lights/tplink"
)

// Scan implements a searcher for local network light devices
func Scan(ip string) (*nmap.Run, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts(tplink.TPLINK_API_PORT),
		nmap.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}

	result, err := scanner.Run()
	if err != nil {
		return nil, err
	}

	return result, nil
}
