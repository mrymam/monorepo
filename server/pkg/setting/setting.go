package setting

type Setting struct {
	Sqlite Sqlite `yaml:"sqlite"`
	JWT    JWT    `yaml:"jwt"`
}
