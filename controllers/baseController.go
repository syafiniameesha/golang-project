package controllers

import (
    "gorm.io/gorm"
)

type BaseController struct {
    DB *gorm.DB
}

func NewBaseController(db *gorm.DB) *BaseController {
    return &BaseController{
        DB: db,
    }
}

func (bc *BaseController) Create(model interface{}) error {
    return bc.DB.Create(model).Error
}

func (bc *BaseController) FindByID(model interface{}, id interface{}) error {
    return bc.DB.First(model, id).Error
}

func (bc *BaseController) FindAll(models interface{}) error {
    return bc.DB.Find(models).Error
}

func (bc *BaseController) Update(model interface{}) error {
    return bc.DB.Save(model).Error
}

func (bc *BaseController) Delete(model interface{}, id interface{}) error {
    return bc.DB.Delete(model, id).Error
}
