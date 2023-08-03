package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	HttpAddress           string `json:"http_address"`
	Port                  string `json:"port"`
	DbSchema              string `json:"db_schema"`
	DbUsername            string `json:"db_username"`
	DbPassword            string `json:"db_pass"`
	DbName                string `json:"db_name"`
	DbAddress             string `json:"db_address"`
	DbPort                string `json:"db_port"`
	DbSslMode             string `json:"db_ssl_mode"`
	RootDownloadPath      string `json:"root_download_path"`
	UserContentUploadPath string `json:"user_content_upload_path"`
	UrlsPath              string `json:"urls_path"`
}

func Load(confPath string) (Config, error) {
	confPath = filepath.Clean(confPath)
	content, err := os.ReadFile(confPath)
	if err != nil {
		return Config{}, err
	}
	conf := new(Config)
	err = json.Unmarshal(content, conf)
	if err != nil {
		return Config{}, err
	}
	return *conf, nil
}
