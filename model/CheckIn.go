package model

type CheckIn struct {
	UserId      int    `json:"userId"`
	TaskId      int    `json:"taskId"`
	CheckedIn   string `json:"checkedIn"`
	CheckInTime string `json:"checkinTime"`
}
