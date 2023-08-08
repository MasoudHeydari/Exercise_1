package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var conf Config

type Config struct {
	HttpAddress           string `json:"http_address"`
	Port                  string `json:"port"`
	DbSchema              string `json:"db_schema"`
	DbUsername            string `json:"db_username"`
	DbPassword            string `json:"db_pass"`
	DbName                string `json:"db_name"`
	DbAddress             string `json:"db_address"`
	DbSslMode             string `json:"db_ssl_mode"`
	RootDownloadPath      string `json:"root_download_path"`
	MaxImageSizeInBytes   int64  `json:"max_image_size_in_bytes"`
	UserContentUploadPath string `json:"user_content_upload_path"`
	UrlsPath              string `json:"urls_path"`
}

func SetConfig(config Config) {
	conf = config
}

// Load loads the config from config.json file and creates a new config instance.
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

func GetServerAddressAndPort() string {
	fmt.Println(conf)
	return fmt.Sprintf("%s:%s", conf.HttpAddress, conf.Port)
}

func GetRootDownloadPath() string {
	return conf.RootDownloadPath
}

func GetUserContentUploadPath() string {
	return conf.UserContentUploadPath
}

func GetMaxImageSizeInBytes() int64 {
	return conf.MaxImageSizeInBytes
}
