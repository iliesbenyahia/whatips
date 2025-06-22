package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetPublicIp() (string, error) {
	type ipifyData struct {
		IP string `json:"ip"`
	}
	var ipifyResp ipifyData

	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &ipifyResp)
	if err != nil {
		fmt.Println("Erreur:", err)
		return "", err
	}
	return ipifyResp.IP, err
}

func GetSavedIP(path string) (string, error) {
	if _, err := os.Stat(path); err == nil {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Erreur lecture :", err)
			return "", err
		}
		return string(data), err
	} else {
		return "", err
	}
}

func SaveIP(ip string, path string) {
	os.WriteFile(path, []byte(ip), 0644)
}
