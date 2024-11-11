package controller

import (
	"net/http"

	"schedule-api-specification/model"

	"github.com/gin-gonic/gin"
)

// GetSchedule retreves the schedule for a specific date.
// @Summary Get schedule by date
// @Description Get a schedule by date
// @Param date path string true "Date"
// @Success 200 {object} model.TaskDTO[]
// @Router /schedule/{date} [get]
func GetSchedule(c *gin.Context) {
	// date := c.Param("date")
	// id := c.Request.BasicAuth("id")
	// // Get the schedule
	// schedule := getTasks(id, date)
	schedule := []model.TaskDTO{
		{Id: 1, GroupId: 1, Name: "Task 1", Description: "Description 1", Date: "2019-01-01", StartTime: "09:00", EndTime: "10:00", Location: "Location 1"},
		{Id: 2, GroupId: 1, Name: "Task 2", Description: "Description 2", Date: "2019-01-01", StartTime: "10:00", EndTime: "11:00", Location: "Location 2"},
		{Id: 3, GroupId: 1, Name: "Task 3", Description: "Description 3", Date: "2019-01-01", StartTime: "11:00", EndTime: "12:00", Location: "Location 3"},
		{Id: 4, GroupId: 1, Name: "Task 4", Description: "Description 4", Date: "2019-01-01", StartTime: "12:00", EndTime: "13:00", Location: "Location 4"},
	}
	// // Return the schedule
	c.JSON(http.StatusOK, schedule)
}

// GetTasks retreves the tasks for a specific date.
// @Summary Get the tasks by date and group
// @Description Get a list of tasks by date and group
// @Param date path string true "Date"
// @Param groupid path string true "Group ID"
// @Success 200 {object} model.TaskDTO[]
// @Router /schedule/{date}/{groupid} [get]
func GetTasks(c *gin.Context) {
	// date := c.Param("date")
	// // Get the tasks for the date and group
	// schedule := getTasks(id, date)
	schedule := []model.TaskDTO{
		{Id: 1, GroupId: 1, Name: "Task 1", Description: "Description 1", Date: "2019-01-01", StartTime: "09:00", EndTime: "10:00", Location: "Location 1"},
		{Id: 2, GroupId: 1, Name: "Task 2", Description: "Description 2", Date: "2019-01-01", StartTime: "10:00", EndTime: "11:00", Location: "Location 2"},
		{Id: 3, GroupId: 1, Name: "Task 3", Description: "Description 3", Date: "2019-01-01", StartTime: "11:00", EndTime: "12:00", Location: "Location 3"},
		{Id: 4, GroupId: 1, Name: "Task 4", Description: "Description 4", Date: "2019-01-01", StartTime: "12:00", EndTime: "13:00", Location: "Location 4"},
	}
	// // Return the schedule
	c.JSON(http.StatusOK, schedule)
}

// GetTask retreves a specific Task.
// @Summary Get a task by ID
// @Description Get a task by ID
// @Param id path string true "ID"
// @Success 200 {object} model.TaskDTO
// @Router /schedule/task/{id} [get]
func GetTask(c *gin.Context) {

	// // Return the task
	c.JSON(http.StatusOK, model.TaskDTO{Id: 1, GroupId: 1, Name: "Task 1", Description: "Description 1", Date: "2019-01-01", StartTime: "09:00", EndTime: "10:00", Location: "Location 1"})
}

// CreateTask creates a new Task.
// @Summary Create a new task
// @Description Create a new task
// @Param task body model.Task true "Task"
// @Success 200 {object} model.TaskDTO
// @Router /schedule/task [post]
func CreateTask(c *gin.Context) {
	var task model.TaskDTO
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Here you would add code to save the task to a database.
	c.JSON(http.StatusOK, task)
}

// UpdateTask creates a Task.
// @Summary Updates a task
// @Description Update a task
// @Param id path string true "ID"
// @Param task body model.Task true "Task"
// @Success 200 {object} model.TaskDTO
// @Router /schedule/task/{id} [put]
func UpdateTask(c *gin.Context) {
	// (int)id := c.Param("id")
	var task model.TaskDTO
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// task.Id = id
	// Here you would add code to save the task to a database.
	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a Task.
// @Summary Deletes a task
// @Description Delete a task
// @Param id path string true "ID"
// @Success 200 {object} model.TaskDTO
// @Router /schedule/task/{id} [delete]
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	// Here you would delete the task from the database using the id.
	c.JSON(http.StatusOK, gin.H{"message": "Task with ID " + id + " deleted successfully"})
}

// CheckIn checks in on a Task.
// @Summary CheckIn on a task
// @Description CheckIn on a task
// @Param id path string true "ID"
// @Success 200 {object} model.CheckInDTO
// @Router /schedule/task/{id} [post]
func CheckIn(c *gin.Context) {
	// taskId := c.Param("id")
	// userId :=

	var checkin model.CheckInDTO
	// Here you would upload the checkin to the database.
	c.JSON(http.StatusOK, checkin)
}

func getTasks(id int, date string) {
	// tasks := []model.Task{
	// 	{1, 1, "Task 1", "Description 1", "2019-01-01", "09:00", "10:00", "Location 1"},
	// 	{2, 1, "Task 2", "Description 2", "2019-01-01", "10:00", "11:00", "Location 2"},
	// 	{3, 1, "Task 3", "Description 3", "2019-01-01", "11:00", "12:00", "Location 3"},
	// 	{4, 1, "Task 4", "Description 4", "2019-01-01", "12:00", "13:00", "Location 4"},
	// }
	// return tasks
}
