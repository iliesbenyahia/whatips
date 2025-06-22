package main

import (
	"fmt"
	"log"
	"os"

	gandi "github.com/iliesbenyahia/whatips/registrar"
	"github.com/iliesbenyahia/whatips/utils"
)

func main() {
	personnalAccessToken := os.Getenv("GANDI_PAT")
	fqdn := os.Getenv("FQDN")
	recordType := os.Getenv("RECORD_TYPE")
	recordName := os.Getenv("RECORD_NAME")

	publicIP, err := utils.GetPublicIp()
	if err != nil {
		log.Fatalf("Failed to get public IP: %v", err)
	}
	fmt.Printf("Public IP is: %s\n", publicIP)

	record, err := gandi.GetDnsRecordByNameAndType(personnalAccessToken, fqdn, recordType, recordName)
	if err != nil {
		log.Fatalf("Failed to get DNS record: %v", err)
	}

	if len(record.RRSetValues) == 0 {
		log.Fatalf("DNS record exists but has no values!")
	}

	currentIP := record.RRSetValues[0]
	fmt.Printf("Registered IP for %s (%s %s) is: %s\n", fqdn, recordType, recordName, currentIP)

	if publicIP != currentIP {
		fmt.Println("Public IP has changed. Updating DNS record...")
		record.RRSetValues[0] = publicIP

		err = gandi.UpdateDnsRecord(personnalAccessToken, fqdn, recordType, recordName, record)
		if err != nil {
			log.Fatalf("Failed to update DNS record: %v", err)
		}
		fmt.Println("DNS record updated successfully.")
	} else {
		fmt.Println("No update needed. DNS record is up to date.")
	}
}
