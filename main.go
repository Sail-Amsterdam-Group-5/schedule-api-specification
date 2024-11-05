package schedule_api_specification

import (
	"net/http"
)

func main() {
	// Register the controller
	http.HandleFunc("/schedule", scheduleController)
	// Start the server
	http.ListenAndServe(":8080", nil)
}
