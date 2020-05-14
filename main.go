//go-aws-dyndns
//ported from python to golang to learn
//https://github.com/danpilch/aws-dyndns

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	//	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"io/ioutil"
	"net/http"
)

var publicIpService = "https://httpbin.org/ip"

func main() {
	// Get environment variables
	hostedZoneId := os.Getenv("AWS_HOSTED_ZONE_ID");
	if hostedZoneId == "" {
		panic("missing AWS_HOSTED_ZONE_ID environment variable")
	}

	// Make https request to get public IP
	resp, err := http.Get(publicIpService)
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
	// Check if multiple comma separated IP addresses are returned
	// if so, return first element 
	var publicIp string = dat["origin"].(string)
	if strings.Contains(publicIp, ",") {
		publicIp = strings.Split(publicIp, ",")[0]
	}
	fmt.Println(publicIp)
	// define aws sdk session
	AwsSession, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	// Create route53 service
	svc := route53.New(AwsSession)
	// Find ListResourceRecordSets
	//recordSets = svc.ListResourceRecordSets()
	// Check if IP already exists 
	fmt.Println(svc)
	fmt.Println(hostedZoneId)
	return
}
