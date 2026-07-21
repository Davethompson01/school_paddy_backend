package students

import "time"

type Project struct {
	UserID       int       `json:"user_id"`
	Category     string    `json:"category" validate:"required"`
	Level        string    `json:"level" validate:"required"`
	Topic        string    `json:"topic" validate:"required,min=5"`
	Description  string    `json:"description" validate:"required,min=20"`
	BidAmount    int       `json:"bidAmount" validate:"required,gt=0"`
	Deadline     time.Time `json:"deadline" validate:"required"`
	UpdatedAt    time.Time `json:"updated_at"`
	Requirement  string    `json:"requirement" validate:"required,min=20"`
	DiscountCode string    `json:"discount_code"`
}

type FileDispute struct {
	UserID             int    `json:"user_id"`
	Solution_expert_id int    `json:"student_id"`
	DisputeCategory    string `json:"dsipute_category"`
	Home_Work_File     string `json:"homework"`
	Dispute_details    string `json:"dispute_details"`
}

type AcceptBid_HomeWork struct {
	Student_id         int `validate:"required"`
	Solution_expert_id int `validate:"required"`
	Project_id int  `validate:"required"`
	Accepted   bool `validate:"required"`
}
