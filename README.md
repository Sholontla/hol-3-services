# hol-3-services

## You see the diagram impmention: 
https://github.com/Sholontla/hol-3-services/blob/master/hol-3-service-diagram.pdf

In this "Application Integration Hands On Labs" developes the HOL-3 hands ON Lab

## HOL-3 - API Trigger to BigQuery - 
  HOL-3-1 : Part 1 - Create an API that can expose a BigQuery table insert row operation
  HOL-3-2 : Part 2 - Change the communication to PubSub instead of just calling the integration directly 

### Part 1:
  The API consists in one endPoint 
  
  localhost:8080/insert
   
  paylaod for HTTP/1.0 client
{
  "column1": string,
  "column2": int,
  "column3": float,
  "column4": string
}

the code will generate
uuid.New for the ID
CreatedAt time.Now().Format("2006-01-02T15:04:05") for the time created

the schema to insert into BigQuery will look like 

{
  "insertID": uuid.New().String()
  "column1": string,
  "column2": int,
  "column3": float,
  "column4": string
  "created_at": time.Now().Format("2006-01-02T15:04:05") 
}

### Part 2:
localhost:8080/pub

Will trigger the ublisher service in PUB/SUB GCP
and save the data into bigQuery GCP


