package session

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
)

func StartNewSession() *session.Session {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
	}
	return sess
}

func SetConfiguration(awsKeyID string, awsSecretKey string, awsRegion string) {
	os.Setenv("AWS_REGION", awsRegion)
	os.Setenv("AWS_ACCESS_KEY_ID", awsKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", awsSecretKey)
}
