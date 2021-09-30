package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	// CommonRewardPool represents configs of the common reward pool's node.
	CommonRewardPool struct {
		ID string `json:"id"`
	}
)

// Read reads config yaml file from path.
func (p *CommonRewardPool) Read(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) { _ = f.Close() }(f)

	decoder := yaml.NewDecoder(f)

	return decoder.Decode(p)
}
