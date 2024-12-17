package dto

import (
	"tunnel-hel/model"
)

type ProxyInfoDto struct {
	ProxyDto
	Status     int8   `json:"status"`
	CreateTime int64  `json:"create_time"`
	ActiveTime int64  `json:"active_time"`
	Country    string `json:"country"`
	Region     string `json:"region"`
	City       string `json:"city"`
	Isp        string `json:"isp"`
}

type ProxyDto struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Proto    string `json:"proto"`
	User     string `json:"user"`
	Password string `json:"password"`
	Source   string `json:"source"`
	Ping     int64  `json:"ping"`
}

func NewProxyDto(m model.ProxyModel) ProxyInfoDto {
	return ProxyInfoDto{
		ProxyDto: ProxyDto{
			Host:     m.Host,
			Port:     m.Port,
			Proto:    m.Proto,
			User:     m.User,
			Password: m.Password,
			Source:   m.Source,
			Ping:     m.Ping,
		},
		Status:     m.Status,
		CreateTime: m.CreateTime,
		ActiveTime: m.ActiveTime,
		Country:    m.Country,
		Region:     m.Region,
		City:       m.City,
		Isp:        m.Isp,
	}
}
