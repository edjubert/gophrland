package cmd

type FloatOptions struct {
	Offset          float64 `yaml:"offset,omitempty"`
	RandomizeCenter bool    `yaml:"randomize_center,omitempty"`
}

func LoadPlugin() {}
