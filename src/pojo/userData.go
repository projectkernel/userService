package pojo

type UserData struct {
	Name string `json:"name"`
	Email string `json:"email"`
	ImageURL string `json:"imageURL"`
	Language string `json:"language"`
	Location string `json:"location"`
	Provider string `json:"provider"`
}