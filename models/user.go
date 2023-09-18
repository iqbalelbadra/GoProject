package models

type User struct {
    ID_User  uint   `gorm:"primaryKey autoIncrement" json:"id_user"`
    Email    string `gorm:"unique" json:"email"`
    Password string `gorm:"not null" json:"-"`
    Nama     string `json:"nama"`
    NoHP     string `json:"no_hp"`
    Alamat   string `json:"alamat"`
}
