package cmd

type Parallax struct{}
type ParallaxOptions struct {
	Name string `yaml:"name"`
}

const (
	Name = "parallax"
)

func LoadPlugin() {}
