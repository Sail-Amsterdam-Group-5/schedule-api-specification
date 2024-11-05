package controller

import (
	"encoding/json"
	"net/http"
)

func getSchedule(w http.ResponceWriter, id int, date string) {
	// Get the schedule
	schedule := getTasks(id, date)
	// Return the schedule
	encoder := json.NewEncoder(w).Encode(schedule)
	return encoder
}

func getTasks(id int, date string) []Task {
	tasks := []Task{
		{1, 1, "Task 1", "Description 1", "2019-01-01", "09:00", "10:00", "Location 1"},
		{2, 1, "Task 2", "Description 2", "2019-01-01", "10:00", "11:00", "Location 2"},
		{3, 1, "Task 3", "Description 3", "2019-01-01", "11:00", "12:00", "Location 3"},
		{4, 1, "Task 4", "Description 4", "2019-01-01", "12:00", "13:00", "Location 4"},
	}
	return tasks
}

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
