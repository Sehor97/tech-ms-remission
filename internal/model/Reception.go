package model

import "time"

type Reception struct {
	Reason        string `json:"reason,omitempty"`
	Observation   string `json:"observation,omitempty"`
	Customer      int    `json:"customer,omitempty"`
	Equipment     int    `json:"equipment,omitempty"`
	Status        Status `json:"status,omitempty"`
	dateAdmission time.Time
	dateDelivery  time.Time
	createAt      time.Time
	updateAt      time.Time
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type Status string

const (
	Pending    Status = "Pending"
	Started    Status = "Started"
	InProgress Status = "InProgress"
	Finished   Status = "Finished"
	Failed     Status = "Failed"
)
