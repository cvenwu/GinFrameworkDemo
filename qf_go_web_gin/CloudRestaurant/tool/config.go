package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

/**
 * @Author: yirufeng
 * @Email: yirufeng@foxmail.com
 * @Date: 2020/10/11 11:14 上午
 * @Desc:
 */

type Config struct {
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}

var _cfg *Config

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&_cfg)
	if err != nil {
		return nil, err
	}

	return _cfg, nil
}
