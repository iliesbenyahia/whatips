package main

import (
	"fmt"
	"os"

	"github.com/iliesbenyahia/whatips/utils"
)

func main() {

	var savePath string = "lastip.txt"
	ip, err := utils.GetIp()
	savedIp, _ := utils.GetSavedIP(savePath)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	if ip != savedIp {
		fmt.Printf("Saving IP %s ... \n", ip)
		utils.SaveIP(ip, savePath)
	} else {
		fmt.Printf("IP (%s) has not changed since last time. \n", ip)
	}

	fmt.Printf("Done ! \n")
	os.Exit(0)

}
