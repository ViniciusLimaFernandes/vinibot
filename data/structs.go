package data

type BotSettings struct {
	Bot struct {
		Auth   string `yaml:"auth"`
		Prefix string `yaml:"prefix"`
	}
}
