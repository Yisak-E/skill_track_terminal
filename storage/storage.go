package storage

// imports
import (
	"encoding/json"
	"errors"
	"github.com/Yisak-E/skill_track_terminal/skill"
	"os"
	"path/filepath"
)

const (
	dirName  = ".myskill"
	fileName = "skills.json"
)

func filePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, dirName, fileName), nil
}

func Load() (map[string]skill.Skill, error) {
	path, err := filePath()

	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)

	if errors.Is(err, os.ErrNotExist) {
		return make(map[string]skill.Skill), nil
	}

	if err != nil {
		return nil, err
	}

	skills := make(map[string]skill.Skill)

	if len(data) == 0 {
		return skills, nil
	}

	if err := json.Unmarshal(data, &skills); err != nil {
		return nil, err
	}

	return skills, nil

}

func Save(skills map[string]skill.Skill) error {

	path, err := filePath()

	if err != nil {
		return err
	}

	// create folder if not exist
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	// pretty json
	data, err := json.MarshalIndent(skills, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)

}
