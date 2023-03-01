package model

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`  // active/closed
}
