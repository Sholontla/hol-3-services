package main

import "restservice/internal/service/bigquery_service"

func main() {
	service := bigquery_service.BigQueryService{}
	service.Endpoints()
}
