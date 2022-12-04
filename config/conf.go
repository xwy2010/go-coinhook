package config

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	cfg     *TomlConfig
	once    sync.Once
	cfgLock sync.RWMutex
)

// TomlConfig toml配置
type TomlConfig struct {
	DB       mysql      `toml:"mysql"`
	SRV      server     `toml:"server"`
	Cookies  cookieInfo `toml:"cookie"`
	TeleGram tg         `toml:"telegram"`
}

// Mysql mysql数据库配置
type mysql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	DBname   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Charset  string `toml:"charset"`
	ConnMax  int    `toml:"connection_max"`
	Enabled  bool
}

// Server 服务器配置
type server struct {
	IP   string
	Port string
}

// CookieInfo cookie配置
type cookieInfo struct {
	Domain string
	Name   string
}

// tg 机器人配置
type tg struct {
	Token  string
	Chatid int64
}

// Config 读取配置文件
func Config() *TomlConfig {
	once.Do(ReloadConfig)
	// ReloadConfig()
	cfgLock.RLock()
	defer cfgLock.RUnlock()
	return cfg
}

// ReloadConfig 重新加载配置文件
func ReloadConfig() {
	filePath, err := filepath.Abs("config/conf.toml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("parse toml file once. filePath: %s\n", filePath)
	config := new(TomlConfig)
	if _, err := toml.DecodeFile(filePath, config); err != nil {
		panic(err)
	}
	cfgLock.Lock()
	defer cfgLock.Unlock()
	cfg = config
}
