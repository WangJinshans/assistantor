package config

type AssistantConfig struct {
	Mysql MysqlConfig
	Redis RedisConfig
}

type MysqlConfig struct {
	Address  string
	Database string
	UserName string
	Password string
}

type RedisConfig struct {
	Address  string
	DB       int
	Password string
}
