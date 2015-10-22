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
}



func main() {

  fmt.Println("CloudWatch")

  watcher := &CloudWatcher{NameSpace:"AWS/SQS"}

  listMetricsInNameSpace(watcher)

}

func getCloudWatchWithCredentials() *cloudwatch.CloudWatch {
  config := aws.NewConfig().WithRegion("us-west-2")
  client := cloudwatch.New(config)
  return client
}


func listMetric(cw *CloudWatcher) {
  params := &cloudwatch.ListMetricsInput{
  	Dimensions: []*cloudwatch.DimensionFilter{
  		{ // Required
  			Name:  aws.String(cw.DimensionName), // Required
  			Value: aws.String(cw.DimensionValue),
  		},
  		// More values...
  	},
  	//MetricName: aws.String("MetricName"),
  	//Namespace:  aws.String("Namespace"),
  	//NextToken:  aws.String("NextToken"),
  }
  fetchFromParams(params)
}

func listMetricsInNameSpace(cw *CloudWatcher) {
  params := &cloudwatch.ListMetricsInput{
    Dimensions: []*cloudwatch.DimensionFilter{
      //{ // Required
      //	Name:  aws.String("DimensionName"), // Required
      //	Value: aws.String("DimensionValue"),
      //},
      // More values...
    },
    //MetricName: aws.String("MetricName"),
    Namespace:  aws.String(cw.NameSpace),
    //NextToken:  aws.String("NextToken"),
  }
  fetchFromParams(params)
}

func listAllMetrics() {
  params := &cloudwatch.ListMetricsInput{
    Dimensions: []*cloudwatch.DimensionFilter{
      //{ // Required
      //	Name:  aws.String("DimensionName"), // Required
      //	Value: aws.String("DimensionValue"),
      //},
      // More values...
    },
    //MetricName: aws.String("MetricName"),
    //Namespace:  aws.String("Namespace"),
    //NextToken:  aws.String("NextToken"),
  }
  fetchFromParams(params)
}

func fetchFromParams(params *cloudwatch.ListMetricsInput) {
  client := getCloudWatchWithCredentials()

  resp, err := client.ListMetrics(params)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println(resp)
}
