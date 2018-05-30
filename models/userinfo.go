package models

// Userinfo is user card
type Userinfo struct {
	ID       string `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	PhotoURL string `json:"photo_url"`
}

// CreateUserinfo builds new card
func CreateUserinfo(id, username, email, fullName, photoURL string) *Userinfo {
	return &Userinfo{
		ID:       id,
		Username: username,
		Email:    email,
		FullName: fullName,
		PhotoURL: photoURL,
	}
}

func (ui *Userinfo) GetId() string {
	return ui.ID
}

func (ui *Userinfo) GetEmail() string {
	return ui.Email
}

func (ui *Userinfo) GetFullName() string {
	return ui.FullName
}

func (ui *Userinfo) GetUsername() string {
	return ui.Username
}

func (ui *Userinfo) GetPhotoUrl() string {
	return ui.PhotoURL
}
