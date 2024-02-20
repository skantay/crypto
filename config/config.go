package config

type Config struct {
	Database Database
}

type Database struct {
	Postgres Postgres
}

type Postgres struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
	SSLMode  string
}
