package domain

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type IP struct {
	Value    string
	IsActive bool
}

type ServiceDiscovery struct {
	next  HttpClient
	hosts map[string][]IP
}

func NewServiceDiscovery(configFile string, next HttpClient) *ServiceDiscovery {
	config, err := parseConfig(configFile)
	if err != nil {
		panic("parse config error:" + err.Error())
	}

	return &ServiceDiscovery{hosts: config, next: next}
}

func (d *ServiceDiscovery) SendRequest(req *HttpRequest) (*HttpResponse, error) {

	//先檢查 hosts 欄位 有沒有對應的 req.Host 資料(IPs)
	//沒有的話， IPs 就會是 Host 本身
	//同時把該變數的指標存到 request 當中
	if val, exists := d.hosts[req.Host]; !exists {
		req.IPs = []IP{{Value: req.Host, IsActive: true}}
	} else {
		req.IPs = val
	}

	// targetIP 預設為第一個 active ip
	ip := getFirstActiveIP(req.IPs)
	if ip == nil {
		return nil, &NoActiveIpErr{}
	}

	req.TargetIp = ip.Value
	fmt.Println("[Service Discovery] target ip:", req.TargetIp)
	res, err := d.next.SendRequest(req)

	// request 發送失敗
	if err != nil {
		fmt.Println("[Service Discovery] request failed:", err.Error())

		//將 hosts[req.Host] 當中符合 target_ip == req.target_ip 的元素刪除(active = false)
		d.deactivateIP(req.Host, req.TargetIp)
		fmt.Printf("[Service Discovery] after removed %s, ips: %v\n", req.TargetIp, d.hosts[req.Host])

		//並且添加一個倒數計時器，當 10 分鐘一到，就把 target_ip 加回去
		time.AfterFunc(10*time.Second, func() {
			fmt.Println("[Service Discovery] add back:", req.TargetIp)
			d.activateIP(req.Host, req.TargetIp)
			fmt.Println("[Service Discovery] ips:", d.hosts[req.Host])
		})
	}

	return res, err
}

func getFirstActiveIP(ips []IP) *IP {
	for _, ip := range ips {
		if ip.IsActive {
			return &ip
		}
	}
	return nil
}

func parseConfig(filePath string) (map[string][]IP, error) {
	config := make(map[string][]IP)

	// 開啟檔案
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			host := strings.TrimSpace(parts[0])
			ips := make([]IP, 0)
			strIps := strings.Split(strings.TrimSpace(parts[1]), ",")
			for _, ip := range strIps {
				ips = append(ips, IP{Value: strings.TrimSpace(ip), IsActive: true})
			}
			config[host] = ips
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func (d *ServiceDiscovery) updateIPStatus(host string, ip string, isActive bool) {
	ips, ok := d.hosts[host]
	if !ok {
		return
	}

	for i := range ips {
		if ips[i].Value == ip {
			ips[i].IsActive = isActive
			break
		}
	}
}

func (d *ServiceDiscovery) deactivateIP(host string, ip string) {
	d.updateIPStatus(host, ip, false)
}

func (d *ServiceDiscovery) activateIP(host string, ip string) {
	d.updateIPStatus(host, ip, true)
}
