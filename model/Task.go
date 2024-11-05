package model

type Task struct {
	Id          int    `json:"id"`
	GroupId     int    `json:"groupId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Location    string `json:"location"`
}
