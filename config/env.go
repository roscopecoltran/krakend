package config

type EnvConfig struct {
	Files []string          `json:"files,omitempty" yaml:"files,omitempty" toml:"files,omitempty"`
	Keys  map[string]string `json:"variables,omitempty" yaml:"variables,omitempty" toml:"variables,omitempty"`
}
