package kind

type User struct {
	Name string `json:"name"`
	Email string `json:"email" bson:"_id,omitempty"`
	ImageURL string `json:"imageURL"`
	Language string `json:"language"`
	Location string `json:"location"`
	Provider string `json:"provider"`
	RefreshToken string `json:"-"`
}