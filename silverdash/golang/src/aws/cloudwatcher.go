package aws

import (
    "fmt"
    "github.com/aws/aws-sdk-go/service/cloudwatch"
    "github.com/aws/aws-sdk-go/aws"
)



type CloudWatcher struct {
    DimensionName string
    DimensionValue string
    MetricName string
    NameSpace string

    CloudWatchClient *cloudwatch.CloudWatch
}

func getCloudWatchWithCredentials() *cloudwatch.CloudWatch {
  config := aws.NewConfig().WithRegion("us-west-2")
  client := cloudwatch.New(config)
  return client
}


func (cw *CloudWatcher) ListMetrics() (out string, err error) {
  params := &cloudwatch.ListMetricsInput{}

  if (len(cw.DimensionName) > 0 && len(cw.DimensionName) > 0) {
    params.Dimensions = []*cloudwatch.DimensionFilter{
  		{ // Required
  			Name:  aws.String(cw.DimensionName), // Required
  			Value: aws.String(cw.DimensionValue),
  		},
    }
  }
  if (len(cw.MetricName) > 0 ) {
    params.MetricName = aws.String(cw.MetricName)
  }
  if (len(cw.NameSpace) > 0 ) {
    params.Namespace = aws.String(cw.NameSpace)
  }
  return cw.fetchFromParams(params)
}

func (cw *CloudWatcher) fetchFromParams(params *cloudwatch.ListMetricsInput) (out string, err error) {
  client := getCloudWatchWithCredentials()

  resp, err := client.ListMetrics(params)

  if err != nil {
    fmt.Println(err.Error())
    return "", err
  }

  return resp.String(), nil
}
