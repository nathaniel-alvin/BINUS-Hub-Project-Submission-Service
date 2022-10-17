package model

import (
	"time"
)

type Project struct {
	Projectid    string    `json:"project-id",validate:"min=1"`
	Title        string    `json:"assignment-id",validate:"min=1"`
	Membersid    []string  `json:"members-id"`
	Score        int64     `json:"score"`
	Instructorid string    `json:"instructor-id",validate:"min=1"`
	Coursecode   string    `json:"course-code",validate:"min=1"`
	Dsubmit      time.Time `json:"date-submitted",validate:"min=1"`
	File         string    `json:"file",validate:"min=1"`
	Thumbnail    string    `json:"thumbnail",validate:"min=1"`
}
