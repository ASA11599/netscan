package scanning

import (
	"net"
	"strconv"
	"time"
)

type ScanInfo struct {
	Network string
	Host string
	Port int
}

type ScanResult struct {
	Info ScanInfo
	Open bool
}

func ScanPort(info ScanInfo) ScanResult {
	res := ScanResult{
		Info: info,
		Open: false,
	}
	address := info.Host + ":" + strconv.Itoa(info.Port)
	c, err := net.DialTimeout(info.Network, address, 5 * time.Second)
	if err != nil {
		res.Open = false
	} else {
		defer c.Close()
		res.Open = true
	}
	return res
}

func ScanPortAsync(info ScanInfo) <-chan ScanResult {
	c := make(chan ScanResult)
	go func(ch chan<- ScanResult, info ScanInfo) {
		c <- ScanPort(info)
	}(c, info)
	return c
}

func ScanAllConcurrent(scans []ScanInfo) []ScanResult {
	results := make([]ScanResult, 0)
	scanResults := make(map[ScanInfo]<-chan ScanResult)
	for _, s := range scans {
		scanResults[s] = ScanPortAsync(s)
	}
	for _, r := range scanResults {
		results = append(results, <-r)
	}
	return results
}

func ScanAllSequential(scans []ScanInfo) []ScanResult {
	results := make([]ScanResult, 0)
	scanResults := make(map[ScanInfo]ScanResult)
	for _, s := range scans {
		scanResults[s] = ScanPort(s)
	}
	for _, r := range scanResults {
		results = append(results, r)
	}
	return results
}
