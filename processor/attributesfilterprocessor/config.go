package attributesfilterprocessor

type Config struct {
	Drop []DropAction `mapstructure:"drop" json:"drop" yaml:"drop"`
}

type DropAction struct {
	Key   string `mapstructure:"key" json:"key" yaml:"key"`
	Value string `mapstructure:"value" json:"value" yaml:"value"`
}
