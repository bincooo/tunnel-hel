package model

import "fmt"

//import (
//	"github.com/jinzhu/gorm"
//)
type ProxyModel struct {
	Id         int    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Host       string `gorm:"column:host"`
	Port       string `gorm:"column:port"`
	Status     int8   `gorm:"column:status"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
	ActiveTime int64  `gorm:"column:active_time"`
	Country    string `gorm:"column:country"`
	Region     string `gorm:"column:region"`
	City       string `gorm:"column:city"`
	Isp        string `gorm:"column:isp"`
	CheckCount int64  `gorm:"column:check_count"`
	Source     string `gorm:"column:source"`
	Proto      string `gorm:"proto"`
	User       string `gorm:"user"`
	Password   string `gorm:"password"`
	Ping       int64  `gorm:"ping"`
}

func (ProxyModel) TableName() string {
	return "proxy"
}

func (m ProxyModel) String() string {
	return fmt.Sprintf("%#v", m)
}
