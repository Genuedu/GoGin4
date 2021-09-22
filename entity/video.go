package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1"`
	Email     string `json:"email" validate:"required"`
}

type Video struct {
	Title       string `json:"title" validate:"is-cool"`
	Description string `json:"description"`
	URL         string `json:"url" binding:"required"`
	Author      Person `json:"author" binding:"required"`
}
