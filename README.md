# gcloud_custommetric_writer

* This repo helps you to create custommetric in gcloud which can be used for scaling up the pods.Basically code is return in Golang.

## Creation of custom metric
* Things that you need to be aware is the gcloud_project id and cluster where you want to use it.Please follow the specific commands below
```bash
go build custommetric_creation.go
go run custom_creation.go ${Cluster_name} ${Project_id}

```

##Note
* You can change the metric type to be anything, basically based on your way of work.
