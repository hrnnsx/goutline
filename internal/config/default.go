package config

func ApplyDefaults(config *Config) {
	if config.API.Protocol == "" {
		config.API.Protocol = "rest"
	}
	if config.API.Serialization == "" {
		config.API.Serialization = "json"
	}
}
