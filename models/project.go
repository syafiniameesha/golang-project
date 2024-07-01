package models

import (
    "time"
    "gorm.io/gorm"
)

type Project struct {
    gorm.Model
    ID          uint      `json:"ID" gorm:"primary_key;auto_increment:true"`
    Prefix      string    `gorm:"size:255;not null;" json:"prefix"`
    Suffix      string    `gorm:"size:255;not null;" json:"suffix"`
    FullCode    string    `gorm:"size:255;not null;" json:"fullCode"`
    Name        string    `gorm:"size:255;not null;" json:"name"`
    StartDate   time.Time `json:"startDate"`
    EndDate     time.Time `json:"endDate"`
    StatusID    uint      `json:"statusId"` // Foreign key to Status
    Status      Status    `gorm:"foreignKey:StatusID"` // Association with Status model
    TypeID      uint      `json:"typeId"` // Foreign key to Type
    Type        Type      `gorm:"foreignKey:TypeID"` // Association with Type model
    UserID      uint      `json:"userId"` // Foreign key to User
    User        User      `gorm:"foreignKey:UserID"` // Association with User model
}
