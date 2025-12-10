package models

import "time"

type Point struct {
    ID          int       `json:"id"`
    UserID      int       `json:"user_id"` //fk ke user
    TotalPoint  int       `json:"total_point"`
    Expired     *time.Time`json:"expired"`

	// Relasi
    User        User      `json:"user"` 
}