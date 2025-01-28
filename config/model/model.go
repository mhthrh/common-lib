package model

type MetaData struct {
	AppName    string `json:"appName"`
	Version    string `json:"version"`
	IsTest     bool   `json:"isTest"`
	ExpireDate string `json:"expireDate"`
}
type Secret struct {
	SecretKey     string `json:"secretKey"`
	TokenDuration int    `json:"tokenDuration"`
}
type DB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
	Driver   string `json:"driver"`
	MinCount int    `json:"minCount"`
	MaxCount int    `json:"maxCount"`
}
