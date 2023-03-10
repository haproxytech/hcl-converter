package configuration

import (
	"os"

	"github.com/hashicorp/hcl"
)

type StorageHCL struct {
	filename string
	cfg      *StorageDataplaneAPIConfiguration
}

func (s *StorageHCL) Load(filename string) error {
	s.filename = filename
	cfg := &StorageDataplaneAPIConfiguration{}

	var hclFile []byte
	var err error
	hclFile, err = os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = hcl.Decode(&cfg, string(hclFile))
	if err != nil {
		return err
	}

	s.cfg = cfg

	return nil
}

func (s *StorageHCL) Get() *StorageDataplaneAPIConfiguration {
	if s.cfg == nil {
		s.cfg = &StorageDataplaneAPIConfiguration{}
	}
	return s.cfg
}
