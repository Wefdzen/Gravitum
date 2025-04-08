package config

type ConfigDatabase struct {
	PG_host     string `yaml:"PG_host" env-default:"localhost"`
	PG_port     string `yaml:"PG_port" env-default:"5432"`
	PG_user     string `yaml:"PG_user" env-default:"postgres"`
	PG_db_name  string `yaml:"PG_db_name" env-default:"postgres"`
	PG_password string `yaml:"PG_password" env-default:"1"`
}

func LaunchConfigFile() ConfigDatabase {
	var cfg ConfigDatabase
	cfg.PG_host = "postgres"
	cfg.PG_port = "5432"
	cfg.PG_user = "postgres"
	cfg.PG_db_name = "postgres"
	cfg.PG_password = "12345"
	return cfg
}
