package tools

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/wjaoss/aws-wrapper/session"
)

func CW(namespace string, metricname string, unit string, vmetric float64, dimName string, dimValue string) {

	sess := session.StartNewSession()

	// Create new cloudwatch client.
	svc := cloudwatch.New(sess)

	_, err := svc.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String(namespace),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String(metricname),
				Unit:       aws.String(unit),
				Value:      aws.Float64(vmetric),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String(dimName),
						Value: aws.String(dimValue),
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println("Error adding metrics:", err.Error())
		return
	}

}
