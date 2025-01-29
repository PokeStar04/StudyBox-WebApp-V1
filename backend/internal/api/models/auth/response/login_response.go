package response

// LoginResponse représente les données renvoyées après une connexion réussie.
type LoginResponse struct {
    Token           string `json:"token"`
    ProfileImage    string `json:"profile_image"`
    IsAuthenticated bool   `json:"isAuthenticated"`
    Message         string `json:"message,omitempty"`
	Data         LoginCacheData `json:"data"`
}

type LoginCacheData struct {

	Id            string  `json:"id"`
	BirthDate     string  `json:"birthDate"`
	City          string  `json:"city"`
	Country       string  `json:"country"`
	Email         string  `json:"email"`
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	NumberStreet  string  `json:"numberStreet"`
	ParrainCode   string `json:"parrainCode"`
	ParrainageCode string  `json:"parrainageCode"`
	Phone         string  `json:"phone"`
	Postcode      string  `json:"postcode"`
	Region        string  `json:"region"`
	Street        string  `json:"street"`
}
