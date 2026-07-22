package students

import "time"

type Notification struct {
	Notification_id    int       `json:"notification_id"`
	Created_at         time.Time `json:"created_at"`
	Seen               bool      `json:"seen"`
	Title              string    `json:"title"`
	Solution_expert_id int       `json:"student_id"`
	Student_id         int       `validate:"required"`
	Project_id         int       `json:"project_id"`
}


type NotificationResponse struct {
	ID int `json:"id"`
	// Name      string    `json:"name"`
	Message   string    `json:"message"`
	Seen      bool      `json:"seen"`
	CreatedAt time.Time `json:"created_at"`
}

type ExpertBidNotification struct {
	NotificationID int
	ExpertName     string
	ProjectTitle   string
	Seen           bool
	CreatedAt      time.Time
}
