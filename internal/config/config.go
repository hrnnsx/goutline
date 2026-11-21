package config

type Config struct {
	Project   ProjectConfig
	Framework FrameworkConfig
	API       APIConfig
	Database  *DatabaseConfig
	Cache     *CacheConfig
}

type ProjectConfig struct {
	Name   string
	Module string
}

type FrameworkConfig struct {
	Provider string
}

type APIConfig struct {
	Protocol      string
	Serialization string
}

type DatabaseConfig struct {
	Provider string
	ORM      string
}

type CacheConfig struct {
	Provider string
}
