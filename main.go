package main

import (
	"fmt"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"os"
)

func main() {

	client, err := scw.NewClient(
		scw.WithEnv(),
		scw.WithDefaultRegion(scw.RegionNlAms),
		scw.WithDefaultZone(scw.ZoneNlAms1),
	)
	if err != nil {
		panic(err)
	}

	lbApi := lb.NewAPI(client)

	loadbalancer, err := lbApi.GetLB(&lb.GetLBRequest{
		LBID: os.Args[1],
	})
	if err != nil {
		panic(err)
	}

	instanceApi := instance.NewAPI(client)

	instanceResp, err := instanceApi.ListServers(&instance.ListServersRequest{
		Tags: loadbalancer.Tags,
	})
	if err != nil {
		panic(err)
	}

	var ips []string

	for _, server := range instanceResp.Servers {
		ips = append(ips, server.PublicIP.Address.String())
	}

	backendResp, err := lbApi.ListBackends(&lb.ListBackendsRequest{
		LBID: loadbalancer.ID,
	})
	if err != nil {
		panic(err)
	}

	_, err = lbApi.SetBackendServers(&lb.SetBackendServersRequest{
		BackendID: backendResp.Backends[0].ID,
		ServerIP:  ips,
	})
	fmt.Println("Done!")
}
