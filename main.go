//go-aws-dyndns
//ported from python to golang to learn
//https://github.com/danpilch/aws-dyndns

package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var publicIpService = "https://httpbin.org/ip"

func main() {
	// Get environment variables
	// not very dry..
	hostedZoneIdEnv := os.Getenv("AWS_HOSTED_ZONE_ID")
	if hostedZoneIdEnv == "" {
		panic("missing AWS_HOSTED_ZONE_ID environment variable")
	}
	domainIdEnv := os.Getenv("AWS_HOSTED_ZONE_DOMAIN_NAME")
	if domainIdEnv == "" {
		panic("missing AWS_HOSTED_ZONE_DOMAIN_NAME environment variable")
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
	// Create route53 service
	svc := route53.New(session.New())
	input := &route53.GetHostedZoneInput{
		Id: aws.String(hostedZoneIdEnv),
	}
	result, err := svc.GetHostedZone(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awser.Error to get the
			// Message from error
			fmt.Println(err.Error())
		}
	}
	if strings.Contains(*result.HostedZone.Name, domainIdEnv) {
		fmt.Println("herre!!")
	}

	// Find ListResourceRecordSets
	//recordSets = svc.ListResourceRecordSets()
	// Check if IP already exists
	return
}
