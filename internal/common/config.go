package common

import "github.com/vrischmann/envconfig"

type Config struct {
	StorageType string `envconfig:"default=ram"`
	DBName      string `envconfig:"default=short_urls,postgres_db"`
	DBPort      string `envconfig:"default=5432,database_port"`
	DBHost      string `envconfig:"default=postgres_container,database_host"`
	DBUserName  string `envconfig:"default=root,postgres_user"`
	DBPassword  string `envconfig:"postgres_password"`
}

var CFG Config

func InitConfig(s string) {
	envconfig.Init(&CFG)
	CFG.StorageType = s
}
