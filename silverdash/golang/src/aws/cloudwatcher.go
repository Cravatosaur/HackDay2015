package aws

import (
    "fmt"
    "time"
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

func (cw *CloudWatcher) FetchMetric() (out *cloudwatch.GetMetricStatisticsOutput, err error) {
  client := getCloudWatchWithCredentials()
  now := time.Now()
  later := now.AddDate(0, 0, -1)

  params := &cloudwatch.GetMetricStatisticsInput{
    //Dimensions : []*cloudwatch.Dimension{
    //  { // Required
    //    Name:  aws.String(cw.DimensionName), // Required
    //    Value: aws.String(cw.DimensionValue),
    //  },
    //},

    EndTime:   &now,     // Required
  	Period:     aws.Int64(120),             // Required
  	StartTime:  &later,     // Required
  	Statistics: []*string{ // Required
  		aws.String("Sum"), // Required
      aws.String("SampleCount"), // Required
  		// More values...
  	},
    MetricName: aws.String(cw.MetricName),
    Namespace: aws.String(cw.NameSpace),
  	//Unit: aws.String("Count"),
  }

  req, resp := client.GetMetricStatisticsRequest(params)
  fmt.Println(req)

  err = req.Send()
  if err != nil {
  	// Print the error, cast err to awserr.Error to get the Code and
  	// Message from an error.
  	fmt.Println(err.Error())
  	return nil, err
  }
  fmt.Println("---")
  fmt.Println(resp)

  return resp, nil
}



func (cw *CloudWatcher) ListMetrics() (*cloudwatch.ListMetricsOutput, error) {
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
  fmt.Println(params)
  return cw.fetchFromParams(params)
}

func (cw *CloudWatcher) fetchFromParams(params *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
  client := getCloudWatchWithCredentials()

  resp, err := client.ListMetrics(params)

  if err != nil {
    fmt.Println(err.Error())
    return nil, err
  }

  return resp, nil
}
