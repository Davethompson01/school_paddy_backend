package solutionexpert_model

import (
	"time"
)

type ApplyForHomeWork struct {
	Solution_expert_id        int    `validate:"required"`
	Paddyproject_id           int    `validate:"required"`
	Student_id                int    `validate:"required"`
	Accepted                  bool   `json:"accepted"`
	IsCompleted               bool   `json:"isCompleted"`
	Accepted_a_expert_already bool   `json:"accepted_a_expert_already"`
	Status                    string `json:"status"`
}

type NegotiateProject struct {
	Solution_expert_id int       `validate:"required"`
	Student_id         int       `validate:"required"`
	Paddyproject_id    int       `validate:"required"`
	Price              int       `json:"price"`
	Deadline           time.Time `json:"deadline"`
	Renegotiate        struct {
		Price    int       `json:"price"`
		Deadline time.Time `json:"deadline"`
	}
	AcceptBidPrice int  `json:"acceptbidprice"`
	Seen           bool `json:"seen"`
}

type HomeWorkProfile struct {
}

type FileDispute struct {
	UserID          int    `validate:"required"`
	Student_id      int    `validate:"required"`
	DisputeCategory string `json:"dsipute_category"`
	Home_Work_File  string `json:"homework"`
	Dispute_details string `validate:"required"`
}

type TotalAcceptProject struct {
	UserID    int    `json:"user_id"`
	Role      string `json:"role"`
	Total     int    `json:"total"`
	Completed int    `json:"completed"`
}

type BidCreatedNotification struct {
	StudentID        int  `json:"student_id"`
	SolutionExpertID int  `json:"solution_expert_id"`
	ProjectID        int  `json:"project_id"`
	Seen             bool `json:"seen"`
	Applied          bool `json:"applied"`
}
