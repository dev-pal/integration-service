package cfg

import (
	"encoding/json"
	"os"
)

func AzDevOps() (*App, error) {
	f, err := os.Open("./pkg/cfg/apps.json")
	defer f.Close()
	if err != nil {
		return nil, err
	}

	var settings map[string]App
	jsonParser := json.NewDecoder(f)
	err = jsonParser.Decode(&settings)
	if err != nil {
		return nil, err
	}

	azDevops := settings["azDevOps"]

	return &azDevops, nil
}
