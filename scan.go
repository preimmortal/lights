package smarthome

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"time"

	"github.com/Ullaakut/nmap"
)

type Scan struct {
	Db *Database
}

// Scan implements a searcher for local network devices
func (s *Scan) Scan(ip string) (*nmap.Run, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts(TPLINK_API_PORT),
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

func (s *Scan) parseLinuxIPRouteShow(output []byte) (net.IP, error) {
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

func (s *Scan) findDefaultRoute() (net.IP, error) {
	routeCmd := exec.Command("ip", "route", "show")
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return s.parseLinuxIPRouteShow(output)
}

func (s *Scan) FindFirstIP() (string, error) {
	log.Print("Scanner Finding First IP Address")
	ip, err := s.findDefaultRoute()
	if err != nil {
		return "", err
	}
	iprange := fmt.Sprintf("%s/24", ip.String())

	hostlist, err := s.Scan(iprange)
	if err != nil {
		return "", err
	}

	for _, host := range hostlist.Hosts {
		target := host.Addresses[0].Addr
		log.Print("Host: ", target)
		for _, port := range host.Ports {
			log.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
			if port.State.State == "open" {
				//TODO: We are assuming that any host with open 9999 port is an available host, we should also confirm the hostname
				// insert into database
				return target, nil
			}
		}
	}
	return "", errors.New("No Valid IP")

}

func (s *Scan) Start(test bool) error {
	if s.Db == nil {
		return errors.New("No Database declared")
	}

	Tp := &Tplink{}

	log.Print("Starting Scanner")
	ip, err := s.findDefaultRoute()
	if err != nil {
		log.Print("Could not find iprange", err)
		return err
	}
	iprange := fmt.Sprintf("%s/24", ip.String())

	for {
		log.Print("Scanning IP Range: ", iprange)
		hostlist, err := s.Scan(iprange)
		if err != nil {
			log.Print("Could not scan ips")
			time.Sleep(time.Minute)
			continue
		}

		for _, host := range hostlist.Hosts {
			target := host.Addresses[0].Addr
			for _, port := range host.Ports {
				if port.State.State == "open" {
					log.Printf("Host: %s - %s - %s\n", target, host.Addresses[0].AddrType, host.Addresses[0].Vendor)
					log.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
					// Check if already in db
					b, err := s.Db.HasIp(target)
					if err != nil {
						log.Print("WARNING: Could not check for IP in database: ", err)
						break
					}
					if !b {
						log.Printf("Inserting %s into db\n", target)
						infoBytes, err := Tp.Send(target, TPLINK_API_INFO)
						if err != nil {
							return err
						}
						var info *TplinkInfo
						err = json.Unmarshal(infoBytes, &info)
						if err != nil {
							return err
						}
						log.Print(info)

						var state string
						switch info.System.GetSysInfo.RelayState {
						case 1:
							state = "on"
						case 0:
							state = "off"
						}

						err = s.Db.Insert(target, info.System.GetSysInfo.DevName, target, info.System.GetSysInfo.Alias, state)
						if err != nil {
							log.Print("WARNING: Could not insert IP in database: ", err)
							break
						}
					} else {
						log.Printf("Skipping %s already db\n", target)
					}
				}
			}
		}
		if test {
			log.Printf("Running in test mode")
			break
		}
		time.Sleep(time.Minute)
	}
	return nil
}
