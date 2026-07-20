package students

import "time"

type profile struct {
	Username    string    `json:"name"`
	Level       string    `json:"level"`
	Created_at  time.Time `json:"created_at"`
	Work_Posted int       `json:"int"`
}
