package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

var dn, hzid string
var ttl int64

func init() {
	dn = os.Getenv("R53_DOMAIN_NAME")
	if dn == "" {
		log.Fatalln("error: R53_DOMAIN_NAME is empty")
	}

	hzid = os.Getenv("R53_HOSTED_ZONE_ID")
	if hzid == "" {
		log.Fatalln("error: R53_HOSTED_ZONE_ID is empty")
	}

	ttl, _ = strconv.ParseInt(os.Getenv("R53_TTL"), 10, 64)
	if ttl == 0 {
		ttl = 300
	}
}

func main() {
	resp, err := http.Get("http://ipecho.net/plain")
	if err != nil {
		log.Fatalln("get error:", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln("read error:", err)
	}

	ip := strings.TrimSpace(string(body))

	log.Println("external ip is:", ip)

	svc := route53.New(nil)

	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(dn),
						Type: aws.String(route53.RRTypeA),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(ip),
							},
						},
						TTL: aws.Int64(ttl),
					},
				},
			},
		},
		HostedZoneId: aws.String(hzid),
	}

	if _, err = svc.ChangeResourceRecordSets(params); err != nil {
		log.Fatalln("aws error:", err.Error())
	}

	log.Println("route53 update was a success.")
	os.Exit(0)
}
