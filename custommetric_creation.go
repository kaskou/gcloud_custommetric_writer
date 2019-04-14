package gcloud_custommetric_writer

import (
	"cloud.google.com/go/monitoring/apiv3"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/genproto/googleapis/api/label"
	"google.golang.org/genproto/googleapis/api/metric"
	_ "google.golang.org/genproto/googleapis/api/metric"
	_ "google.golang.org/genproto/googleapis/api/monitoredres"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
	"log"
	"os"
)

const metricType = "custom.googleapis.com/custom_measurement_value"

var Args []string
func projectResource(projectID string) string {
	return "projects/" + projectID
}
// ------creates a custom metric specified by the metric type.----------------------------------------------------------

func createCustomMetric(projectID, metricType string, clusterName string) error {
	fmt.Printf("Metric Type - %s is created", metricType)
	ctx := context.Background()
	c, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		return err
	}
	md := &metric.MetricDescriptor{
		Name: "Custom Metric",
		Type: metricType,
		Labels: []*label.LabelDescriptor{{
			Key:         "environment",
			ValueType:   label.LabelDescriptor_STRING,
			Description: "An messages/consumers in rabbitmq measurement",
		}},
		MetricKind:  metric.MetricDescriptor_GAUGE,
		ValueType:   metric.MetricDescriptor_DOUBLE,
		Unit:        "%",
		Description: "Messages vs Consumers of the queue.Information is gained by the rabbitmq",
		DisplayName: projectID+ "-"+clusterName+"- Metric",

	}
	req := &monitoringpb.CreateMetricDescriptorRequest{
		Name:             "projects/" + projectID,
		MetricDescriptor: md,
	}
	resp, err := c.CreateMetricDescriptor(ctx, req)
	if err != nil {
		return fmt.Errorf("could not create custom metric: %v", err)
	}

	log.Printf("createCustomMetric: %s\n", formatResource(resp))
	return nil
}

//--------- formatResource marshals a response object as JSON.----------------------------------------------------------
func formatResource(resource interface{}) []byte {
	b, err := json.MarshalIndent(resource, "", "    ")
	if err != nil {
		panic(err)
	}
	return b
}

// -------------------Deletion of custom metric created ----------------------------------------------------------------
func deleteMetric(projectID, metricType string) error {
	fmt.Printf("Metric Type - %s gonna be deleted", metricType)
	ctx := context.Background()
	c, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		return err
	}
	metricResource := "projects/" + projectID + "/metricDescriptors/" + metricType
	req := &monitoringpb.DeleteMetricDescriptorRequest{
		Name: metricResource,
	}
	err = c.DeleteMetricDescriptor(ctx, req)
	if err != nil {
		return fmt.Errorf("could not delete metric: %v", err)
	}
	log.Printf("Deleted metric: %q\n", metricType)
	return nil
}
//----------------------------------------------------------------------------------------------------------------------
func main(){

		createCustomMetric(os.Args[2],metricType,os.Args[1])
		//deleteMetric(os.Args[2], metricType)


}