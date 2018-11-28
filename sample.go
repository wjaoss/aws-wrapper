package main

import (
	"github.com/wjaoss/aws-wrapper/session"
	"github.com/wjaoss/aws-wrapper/tools"
)

func main() {

	awsKeyID := "your ID"
	awsSecretKey := "your Key"
	awsRegion := "us-east-1"

	session.SetConfiguration(awsKeyID, awsSecretKey, awsRegion)
	tools.CW("Testing", "Logsize", "Megabytes", 10.3, "LogFile", "/tmp/me.txt")

}
