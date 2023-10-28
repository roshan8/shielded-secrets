package vars

import (
	"fmt"
	"os"
	"strconv"
)

var (
	AWSAccessKey = ""
	AWSSecretKey = ""
	Port         = 9090
	Regions      = []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "af-south-1",
		"ap-east-1", "ap-south-2", "ap-southeast-3", "ap-southeast-4", "ap-south-1", "ap-northeast-3", "ap-northeast-2",
		"ap-southeast-1", "ap-southeast-2", "ap-northeast-1", "ca-central-1", "eu-central-1", "eu-west-1", "eu-west-2",
		"eu-south-1", "eu-west-3", "eu-south-2", "eu-north-1", "eu-central-2", "il-central-1", "me-south-1", "me-central-1", "sa-east-1"}
)

const (
	Meta      = "meta"
	Data      = "data"
	RegionID  = "regionID"
	SecretID  = "secretID"
	Secret    = "secret"
	AwsClient = "awsClient"
)

func Init() {
	var err error
	AWSAccessKey = os.Getenv("AWS_ACCESS_KEY")
	AWSSecretKey = os.Getenv("AWS_SECRET_KEY")
	PortStr := os.Getenv("PORT")
	Port, err = strconv.Atoi(PortStr)
	if err != nil {
		// refactor this
		fmt.Printf("Error converting PORT to integer: %v", err)
		Port = 9090
		fmt.Println("Setting PORT to default value: 9090")
	}
}
