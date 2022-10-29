package setting

type Setting struct {
	Sqlite  Sqlite  `yaml:"sqlite"`
	JWT     JWT     `yaml:"jwt"`
	Twitter Twitter `yaml:"twitter"`
	Slack   Slack   `yaml:"slack"`
}
