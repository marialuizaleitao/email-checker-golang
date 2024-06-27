package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type DomainInfo struct {
	Domain      string
	HasMX       bool
	HasSPF      bool
	SPFRecord   string
	HasDMARC    bool
	DMARCRecord string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter domain names (one per line) to check:")
	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord")

	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" {
			continue
		}
		info := checkDomain(domain)
		fmt.Printf("%s, %v, %v, %s, %v, %s\n", info.Domain, info.HasMX, info.HasSPF, info.SPFRecord, info.HasDMARC, info.DMARCRecord)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Scanner error:", err)
	}
}

func checkDomain(domain string) DomainInfo {
	info := DomainInfo{
		Domain: domain,
	}

	// Check MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records for domain %s: %v", domain, err)
	}
	info.HasMX = len(mxRecords) > 0

	// Check SPF records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records for domain %s: %v", domain, err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			info.HasSPF = true
			info.SPFRecord = record
			break
		}
	}

	// Check DMARC records
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error looking up TXT records for domain %s: %v", domain, err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			info.HasDMARC = true
			info.DMARCRecord = record
			break
		}
	}

	return info
}
