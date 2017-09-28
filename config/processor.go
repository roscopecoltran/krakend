package config

type ProcessorConfig struct {
	Disabled  bool            `mapstructure:"disabled" default:"false"json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug     bool            `mapstructure:"debug" default:"false"json:"debug" yaml:"debug" toml:"debug"`
	Formatter FormatterConfig `mapstructure:"formatter" json:"formatter" yaml:"formatter" toml:"formatter"`
	Indexer   IndexerConfig   `mapstructure:"indexer" json:"indexer" yaml:"indexer" toml:"indexer"`
	TaskQ     []TaskQConfig   `mapstructure:"taskq" json:"taskq" yaml:"taskq" toml:"taskq"`
}

type IndexerConfig struct {
	Disabled bool          `mapstructure:"disabled" default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool          `mapstructure:"debug" default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Engine   string        `mapstructure:"engine" default:"bleve" json:"engine" yaml:"engine" toml:"engine"`
	Format   string        `mapstructure:"format" default:"json" json:"format" yaml:"format" toml:"format"`
	TaskQ    []TaskQConfig `mapstructure:"taskq" json:"taskq" yaml:"taskq" toml:"taskq"`
}

type FormatterConfig struct {
	Disabled bool          `mapstructure:"disabled" default:"false" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool          `mapstructure:"debug" default:"false" json:"debug" yaml:"debug" toml:"debug"`
	Engine   string        `mapstructure:"engine" default:"mxj" json:"engine" yaml:"engine" toml:"engine"`
	Format   string        `mapstructure:"format" default:"json" json:"format" yaml:"format" toml:"format"`
	TaskQ    []TaskQConfig `mapstructure:"taskq" json:"taskq" yaml:"taskq" toml:"taskq"`
}
