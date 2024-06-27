package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    ID                  uint  `json:"ID" gorm:"primary_key;auto_increment:true"`
    FirstName           string `gorm:"size:255;not null;" json:"firstname"`
    LastName            string `gorm:"size:255;not null;" json:"lastname"`
    Email               string `gorm:"size:100;not null;unique" json:"email"`
    Password            string `gorm:"size:100;not null;" json:"password"`
    Token               string `json:"token"`
    RefreshToken        string `json:"refresh_token"`
    PasswordResetToken  string `gorm:"default:null" json:"passwordResetToken"`
    PasswordResetExpiresAt *time.Time `gorm:"default:null" json:"passwordResetExpiresAt"`
}
