package config

import (
	"fmt"
	"os"

	"github.com/cryring/blog_backend/internal/utils"
	"gopkg.in/yaml.v3"
)

var (
	conf Config
)

type Config struct {
	RootDir  string   `yaml:"root_dir"`
	DBConfig DBConfig `yaml:"db_config"`
}

func (cfg *Config) Validate() error {
	if cfg.RootDir == "" {
		return fmt.Errorf("root_dir is empty")
	}
	ok, err := utils.PathExists(cfg.RootDir)
	if err != nil {
		return fmt.Errorf("invalid root_dir[%s] with error: %v", cfg.RootDir, err)
	}
	if !ok {
		return fmt.Errorf("root_dir[%s] not exist", cfg.RootDir)
	}

	if err := cfg.DBConfig.validate(); err != nil {
		return err
	}
	return nil
}

func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read config file failed: %v", err)
	}

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, fmt.Errorf("unmarshal config file failed: %v", err)
	}

	if err := conf.Validate(); err != nil {
		return nil, fmt.Errorf("validate config failed: %v", err)
	}
	return &conf, nil
}

func GetConfig() *Config {
	return &conf
}

type DBConfig struct {
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (cfg *DBConfig) validate() error {
	if cfg.Address == "" {
		return fmt.Errorf("db address is empty")
	}
	if cfg.User == "" {
		return fmt.Errorf("db user is empty")
	}
	if cfg.Password == "" {
		return fmt.Errorf("db password is empty")
	}
	if cfg.Database == "" {
		return fmt.Errorf("db database is empty")
	}
	return nil
}

func (cfg *DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Address, cfg.Database)
}
