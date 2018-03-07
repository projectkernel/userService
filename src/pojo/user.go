package pojo

type User struct {
	Id string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
	Email string `json:"email"`
	ImageURL string `json:"imageURL"`
	Language string `json:"language"`
	Location string `json:"location"`
	Provider string `json:"provider"`
	RefreshToken string `json:"-"`
}