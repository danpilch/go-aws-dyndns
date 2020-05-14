//go-aws-dyndns
//ported from python to golang to learn go
//https://github.com/danpilch/aws-dyndns

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPublicIp() string {
	// Public service to get your public IP
	url := "https://httpbin.org/ip"
	// Make https request
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Get data from body
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// Parse JSON
	var dat map[string]interface{}
	if err := json.Unmarshal(ip, &dat); err != nil {
		panic(err)
	}
	// return public IP as string
	return dat["origin"].(string)
}
func main() {
	// Find public ip
	var publicIp string
	publicIp = GetPublicIp()
	fmt.Println(publicIp)
}
