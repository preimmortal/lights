package scan

import (
	"testing"
)

func TestScan(t *testing.T) {
	local_ip := "192.168.1.0/24"
	t.Log("Scanning ", local_ip)
	result, err := Scan(local_ip)
	if err != nil {
		t.Fatalf("Could not complete scan: %v", err)
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		t.Logf("Host %q: \n", host.Addresses[0])

		for _, port := range host.Ports {
			t.Logf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	t.Logf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
}
