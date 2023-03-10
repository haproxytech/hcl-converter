package configuration

import (
	"os"

	"github.com/google/renameio"
	"gopkg.in/yaml.v2"
)

type StorageYAML struct {
	filename string
	cfg      *StorageDataplaneAPIConfiguration
}

func (s *StorageYAML) Load(filename string) error {

	s.filename = filename
	cfg := &StorageDataplaneAPIConfiguration{}

	var err error
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return err
	}

	s.cfg = cfg
	return nil
}

func (s *StorageYAML) Set(cfg *StorageDataplaneAPIConfiguration) {
	s.cfg = cfg
}

func (s *StorageYAML) SaveAs(filename string) error {
	data, err := yaml.Marshal(s.cfg)
	if err != nil {
		return err
	}
	return renameio.WriteFile(filename, data, 0o644)
}
