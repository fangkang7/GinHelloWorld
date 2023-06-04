package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string    `json:"app_name"`
	AppMode  string    `json:"app_mode"`
	AppHost  string    `json:"app_host"`
	AppPort  string    `json:"app_port"`
	Sms      SmsConfig `json:"sms"`
	Database OrmConfig `json:"database"`
}

type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	RegionId     string `json:"region_id"`
}

type OrmConfig struct {
	Drive    string `json:"drive"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&_cfg); err != nil {
		return nil, err
	}

	return _cfg, nil
}
