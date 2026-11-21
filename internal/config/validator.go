package config

import (
	"errors"
	"strings"
)

func ValidateConfig(config *Config) error {
	if config == nil {
		return errors.New("There's no config loaded.")
	}
	projectName := strings.TrimSpace(config.Project.Name)
	if projectName == "" {
		return errors.New("Project Name is required")
	}

	projectModule := strings.TrimSpace(config.Project.Module)
	if projectModule == "" {
		return errors.New("Project Module is required")
	}

	framework := strings.ToLower(strings.TrimSpace(config.Framework.Provider))
	switch framework {
	case "":
		return errors.New("Please choose the framework")
	case "gin":
	default:
		return errors.New("We still didn't support that framework...")
	}

	APIprotocol := strings.ToLower(strings.TrimSpace(config.API.Protocol))
	switch APIprotocol {
	case "":
		return errors.New("We need to know what protocol you use")
	case "rest":
	default:
		return errors.New("We still didn't support that protocol...")
	}

	APISerialization := strings.ToLower(strings.TrimSpace(config.API.Serialization))
	switch APISerialization {
	case "":
		return errors.New("We need to know what serialization you use")
	case "json":
	default:
		return errors.New("We still didn't support that serialization...")
	}

	if config.Database != nil {
		dbProvider := strings.ToLower(strings.TrimSpace(config.Database.Provider))
		switch dbProvider {
		case "":
			return errors.New("Please choose a database provider")
		case "mysql":
		case "postgresql":
		default:
			return errors.New("We still didn't support that database...")
		}

		if strings.ToLower(strings.TrimSpace(config.Database.ORM)) != "" {
			dbOrm := strings.ToLower(strings.TrimSpace(config.Database.ORM))
			switch dbOrm {
			case "gorm":
			default:
				return errors.New("We still didn't support that orm...")
			}
		}
	}

	if config.Cache != nil {
		dbCache := strings.ToLower(strings.TrimSpace(config.Cache.Provider))
		switch dbCache {
		case "":
			return errors.New("Please choose a cache provider")
		case "redis":
		default:
			return errors.New("We still didn't support that cache provider...")
		}
	}

	return nil
}
