package gandi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GandiDomainRecord struct {
	RRSetName   string   `json:"rrset_name"`
	RRSetType   string   `json:"rrset_type"`
	RRSetTTL    int      `json:"rrset_ttl"`
	RRSetValues []string `json:"rrset_values"`
	RRSetHref   string   `json:"rrset_href"`
}

func GetDnsRecords(personnalAccessToken string, fqdn string) ([]GandiDomainRecord, error) {
	url := "https://api.gandi.net/v5/livedns/domains/" + fqdn + "/records"
	bearer := "Bearer " + personnalAccessToken

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Request creation error:", err)
		return nil, err
	}
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request execution error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body error:", err)
		return nil, err
	}

	fmt.Println("Raw response body:", string(body)) // debug de la réponse brute

	var data []GandiDomainRecord
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return nil, err
	}

	fmt.Printf("Parsed data: %+v\n", data) // debug des données décodées

	return data, nil
}

func GetIPFromRecords(records []GandiDomainRecord, recordType string, recordName string) string {
	for _, record := range records {
		if record.RRSetType == recordType && record.RRSetName == recordName && len(record.RRSetValues) > 0 {
			return record.RRSetValues[0]
		}
	}
	return ""
}

func GetDnsRecordByNameAndType(personnalAccessToken, fqdn, recordType, recordName string) (GandiDomainRecord, error) {
	url := fmt.Sprintf("https://api.gandi.net/v5/livedns/domains/%s/records/%s/%s", fqdn, recordName, recordType)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GandiDomainRecord{}, fmt.Errorf("Get DNS Record : failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+personnalAccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GandiDomainRecord{}, fmt.Errorf("Get DNS Record : request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // lecture même en erreur pour log utile
		return GandiDomainRecord{}, fmt.Errorf("Get DNS Record : API error: %s - %s", resp.Status, string(body))
	}

	var record GandiDomainRecord
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		return GandiDomainRecord{}, fmt.Errorf("Get DNS Record : failed to decode response: %w", err)
	}

	return record, nil
}

func UpdateDnsRecord(personnalAccessToken, fqdn, recordType, recordName string, newDomainRecord GandiDomainRecord) error {
	url := fmt.Sprintf("https://api.gandi.net/v5/livedns/domains/%s/records/%s/%s", fqdn, recordName, recordType)

	jsonData, err := json.Marshal(newDomainRecord)
	if err != nil {
		return fmt.Errorf("failed to encode domain record to JSON: %w", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+personnalAccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Gandi API error: %s - %s", resp.Status, string(body))
	}

	return nil
}
