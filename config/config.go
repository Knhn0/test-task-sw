package config

type Config struct {
	Env       Env
	Port      int
	Databases Databases `json:"databases"`
}

type Databases struct {
	Postgres Database `json:"postgres"`
}

type Database struct {
	ConnectionString string `json:"connectionString"`
	MigrationPath    string `json:"migrationPath"`
}
