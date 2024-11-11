package model

type CheckInDTO struct {
	CheckInId   int    `json:"checkinId"`
	UserId      int    `json:"userId"`
	TaskId      int    `json:"taskId"`
	CheckedIn   bool   `json:"checkedIn"`
	CheckInTime string `json:"checkinTime"`
}
