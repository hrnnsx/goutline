package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hrnnsx/goutline/internal/config"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	config.ApplyDefaults(cfg)

	if err = config.ValidateConfig(cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Configuration is valid!")
	jsoncfg, _ := json.MarshalIndent(cfg, "", "\t")
	fmt.Println(string(jsoncfg))
}
