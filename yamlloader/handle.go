package yamlloader

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// 讀取ymal設定檔
func LoadYaml(path string, conf interface{}) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return err
	}

	return nil
}
