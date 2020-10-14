package schemas

// UserCreate represents a user to be created
type UserCreate struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

// UserUpdate represents a user to be updated
type UserUpdate struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

// User represents a user to be returned as a response
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}
