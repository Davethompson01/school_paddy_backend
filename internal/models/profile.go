package models

import "time"

type Profile struct {
	User_id            int       `json:"user_id"`
	Username           string    `json:"name"`
	Profile_picture    string    `json:"profile_picture"`
	Level              string    `json:"level"`
	Created_at         time.Time `json:"created_at"`
	Work_Posted        int       `json:"int"`
	Role               string    `json:"role"`
	Brief_infxormation string    `json:"brief_information"`
}

type Profile_expert struct {
	User_id            int       `json:"user_id"`
	Username           string    `json:"name"`
	Profile_picture    string    `json:"profile_picture"`
	Level              string    `json:"level"`
	Created_at         time.Time `json:"created_at"`
	Work_Posted        int       `json:"int"`
	Role               string    `json:"role"`
	Brief_infxormation string    `json:"brief_information"`
	Categories         string
	Resume             string
	Certificate        string
}
