package config

type Config struct {
	Env       Env
	Port      int
	Databases Databases `json:"databases"`
	Hash      Hash      `json:"hash"`
}

type Databases struct {
	Postgres Database `json:"postgres"`
}

type Database struct {
	ConnectionString string `json:"connectionString"`
	MigrationPath    string `json:"migrationPath"`
}

type Hash struct {
	SaltSize           int `json:"saltSize"`
	SecurityIterations int `json:"securityIterations"`
	SecurityKeyLen     int `json:"securityKeyLen"`
}
