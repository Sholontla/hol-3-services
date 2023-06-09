package entity

// Define a struct to represent the data you want to insert
type MyData struct {
	InsertID  string  `json:"insertID" bigquery:"insertID"`
	Column1   string  `json:"column1" bigquery:"column1"`
	Column2   int     `json:"column2" bigquery:"column2"`
	Column3   float64 `json:"column3" bigquery:"column3"`
	Column4   string  `json:"column4" bigquery:"column4"`
	CreatedAt string  `json:"created_at" bigquery:"created_at"`
	// Add more fields as needed
}
