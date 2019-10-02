package scan

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/Ullaakut/nmap"
	"github.com/preimmortal/lights/database"
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

func parseLinuxIPRouteShow(output []byte) (net.IP, error) {
	// Linux '/usr/bin/ip route show' format looks like this:
	// default via 192.168.178.1 dev wlp3s0  metric 303
	// 192.168.178.0/24 dev wlp3s0  proto kernel  scope link  src 192.168.178.76  metric 303
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 3 && fields[0] == "default" {
			ip := net.ParseIP(fields[2])
			if ip != nil {
				return ip, nil
			}
		}
	}

	return nil, errors.New("No gateway found")
}

func FindDefaultRoute() (net.IP, error) {
	routeCmd := exec.Command("ip", "route", "show")
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return parseLinuxIPRouteShow(output)
}

func Start() error {
	// Initialize Database
	if err := database.Init(); err != nil {
		return err
	}

	ip, err := FindDefaultRoute()
	if err != nil {
		return err
	}
	iprange := fmt.Sprintf("%s/24", ip.String())

	hostlist, err := Scan(iprange)
	if err != nil {
		return err
	}

	for _, host := range hostlist.Hosts {
		target := host.Addresses[0].Addr
		//fmt.Println("Host: ", host.Addresses[0].Addr)
		fmt.Println("Host: ", target)
		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
			if port.State.State == "open" {
				//TODO: We are assuming that any host with open 9999 port is an available host, we should also confirm the hostname
				// insert into database
				b, err := database.HasIp(target)
				if err != nil {
					return err
				}
				if b {
					database.Insert("TPLink_Plug", host.Addresses[0].Addr, strconv.FormatUint(uint64(port.ID), 10))
				}
			}
		}
	}

	return nil
}
