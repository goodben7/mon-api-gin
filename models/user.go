package models

type User struct {
	ID    string `json:"id" binding:"required,uuid"` // Doit être un UUID
	Name  string `json:"name" binding:"required,min=3,max=50,not_admin"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"gte=18"` // Greater than or equal to 18
}
