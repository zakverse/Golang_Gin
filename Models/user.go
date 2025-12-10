package models

import (
    "time"
)

type User struct {
    ID           string    	`json:"id"`     
    Username     string    	`json:"username"`        
    Password     string    	`json:"password"`         
    JoinDate     time.Time 	`json:"join_date"`     
    BirthDate    *time.Time	`json:"birth_date"`     
    ReferralCode *string 	`json:"referral_code"`              
}