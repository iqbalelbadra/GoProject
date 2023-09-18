package models

type UserResponse struct {
	ID_User   uint   `json:"id_user"`
    Email    string `json:"email"`
    Nama     string `json:"nama"`
    NoHP     string `json:"no_hp"`
    Alamat   string `json:"alamat"`
	Token	 string `json:"token"`
}