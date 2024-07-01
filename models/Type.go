package models

import (
    "gorm.io/gorm"
)

type Type struct {
    gorm.Model
    ID          uint      `json:"ID" gorm:"primary_key;auto_increment:true"`
    Name        string    `gorm:"size:255;not null;" json:"name"`
}
