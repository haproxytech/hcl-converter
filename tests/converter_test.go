/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/haproxytech/hcl-converter/configuration"
)

func TestConverter(t *testing.T) {
	configs := map[string]bool{
		"../example.hcl":              true,
		"../example-2.hcl":            true,
		"../example-non-existing.hcl": false,
	}

	for config, shouldPass := range configs {

		t.Run(config, func(t *testing.T) {
			storageHCL := &configuration.StorageHCL{}
			err := storageHCL.Load(config)

			if shouldPass {
				if err != nil {
					t.Error("Error while loading the input HCL configuration")
				}

				storageYAML := &configuration.StorageYAML{}
				storageYAML.Set(storageHCL.Get())

				output := strings.Replace(config, filepath.Ext(config), ".yaml", 1)
				err = storageYAML.SaveAs(output)
				if err != nil {
					t.Error(err)
				}
				return
			}
			if err == nil {
				t.Errorf(fmt.Sprintf("error: did not throw error for config file [%s]", config))
			}
		})

		t.Cleanup(func() {
			if _, err := os.Stat("../example.yaml"); err == nil {
				os.Remove("../example.yaml")
			}
			if _, err := os.Stat("../example-2.yaml"); err == nil {
				os.Remove("../example-2.yaml")
			}
		})
	}
}
