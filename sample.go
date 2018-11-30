package main

import (
	"github.com/wjaoss/aws-wrapper/apps"
	"github.com/wjaoss/aws-wrapper/session"
	"github.com/wjaoss/aws-wrapper/tools"
)

func this() {

	awsKeyID := "your-ID"
	awsSecretKey := "your-secret-key"
	awsRegion := "region"

	session.SetConfiguration(awsKeyID, awsSecretKey, awsRegion)
	tools.CW("Testing", "Logsize", "Megabytes", 10.3, "LogFile", "/tmp/me.txt")
	apps.SES("sender@gmail.com", "recipient@gmail.com", "subject", "body")

}
