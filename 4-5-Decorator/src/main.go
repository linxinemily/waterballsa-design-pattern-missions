package main

import (
	"4-5/domain"
	"fmt"
	"time"
)

func main() {

	// 服務探索 → 負載平衡 → 黑名單
	client := domain.NewServiceDiscovery("service_discovery_config",
		domain.NewLoadBalancing(
			domain.NewBlacklist("blacklist_config", domain.NewFakeHttpClient())),
	)

	// 負載平衡 → 服務探索 → 黑名單
	//client := domain.NewLoadBalancing(
	//	domain.NewServiceDiscovery("service_discovery_config",
	//		domain.NewBlacklist("blacklist_config", domain.NewFakeHttpClient())))

	for i := 0; i < 6; i++ {
		response, err := client.SendRequest(domain.NewHttpRequest("http://waterballsa.tw/waterball"))

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("response status:", response.Status)
		}
	}

	for {
		time.Sleep(time.Second)
	}
}
