package models

type Type struct {
    ID          uint      `json:"id"`
    CodeType    int       `json:"code_type"`
    Description string    `json:"description"`
}