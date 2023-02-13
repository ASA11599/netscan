package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ASA11599/netscan/scanning"
	"github.com/ASA11599/netscan/utils"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Error: must pass exactly one CSV file as argument")
	}
	csvFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error: file not found")
	}
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error: could not read records")
	}
	scans := make([]scanning.ScanInfo, 0)
	for _, r := range records {
		if len(r) != 3 {
			continue
		}
		rNetwork := strings.TrimSpace(r[0])
		rHost := strings.TrimSpace(r[1])
		rPort := strings.TrimSpace(r[2])
		if utils.ValidNetwork(rNetwork) && utils.ValidPort(rPort) {
			rPortNumber, err := strconv.Atoi(rPort)
			if err != nil {
				log.Fatal("Error: could not read port number")
			}
			scans = append(scans, scanning.ScanInfo{Network: rNetwork, Host: rHost, Port: rPortNumber})
		}
	}
	for _, s := range scanning.ScanAllConcurrent(scans) {
		msg := "open"
		if !s.Open { msg = "closed" }
		fmt.Printf("%s %s:%d %s\n", s.Info.Network, s.Info.Host, s.Info.Port, msg)
	}
}
