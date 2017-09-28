package config

type EntityConfig struct {
	Disabled bool               `default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool               `default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Mappings []EntityMapsConfig `json:"mappings" yaml:"mappings" toml:"mappings"`
}

type EntityMapsConfig struct {
	Disabled bool   `default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool   `default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Key      string `json:"key" yaml:"key" toml:"key"`
	Map      string `json:"map" yaml:"map" toml:"map"`
	Type     string `json:"type" yaml:"type" toml:"type"`
	Format   string `json:"format" yaml:"format" toml:"format"`
}
