package solutionexpert_model

import (
	"time"
)

type ApplyForHomeWork struct {
	UserID          int  `json:"user_id"`
	Paddyproject_id int  `json:"Paddyproject_id"`
	Student_id      int  `json:"student_id"`
	Accepted        bool `json:"accepted"`
	IsCompleted     bool `json:"isCompleted"`
}

type NegotiateProject struct {
	UserID          int       `json:"user_id"`
	Student_id      int       `json:"student_id"`
	Accepted        bool      `json:"accepted"`
	Paddyproject_id int       `json:"Paddyproject_id"`
	Price           int       `json:"price"`
	Deadline        time.Time `json:"deadline"`
	Renegotiate     struct {
		Price    int       `json:"price"`
		Deadline time.Time `json:"deadline"`
	}
	AcceptBidPrice int `json:"acceptbidprice"`
}

type HomeWorkProfile struct {
}

type FileDispute struct {
	UserID          int    `json:"user_id"`
	Student_id      int    `json:"student_id"`
	DisputeCategory string `json:"dsipute_category"`
	Home_Work_File  string `json:"homework"`
	Dispute_details string `json:"dispute_details"`
}
