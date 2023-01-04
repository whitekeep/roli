package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Token        string `json:"token"`
	GuildId      string `json:"guild_id"`
	ActivityType string `json:"activity_type"`
	ActivityName string `json:"activity_name"`
	Roles        Roles  `json:"roles"`
}

type Roles struct {
	AdminRoles     []string `json:"admin_roles"`
	ModeratorRoles []string `json:"moderator_roles"`
	DevRoles       []string `json:"developer_roles"`
	OwnerRoles     []string `json:"owner_roles"`
}

func GetConfig() (*Config, error) {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		config := new(Config)

		bytes, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return nil, err
		}

		// write json to file
		err = os.WriteFile("config.json", bytes, 0644)
		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile("config.json")

	if err != nil {
		return nil, err
	}

	config := new(Config)

	_ = json.Unmarshal(data, config)

	return config, nil
}
