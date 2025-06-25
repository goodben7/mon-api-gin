package models

type User struct {
	ID    string `json:"id" binding:"required,uuid" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Name  string `json:"name" binding:"required,min=3,max=50" example:"Jean Dupont"`
	Email string `json:"email" binding:"required,email" example:"jean.dupont@example.com"`
	Age   int    `json:"age" binding:"gte=18" example:"25"`
}
