package config

type SenderConf struct {
	Method string  `yaml:"method"`
	Ftp    FtpConf `yaml:"ftp"`
}

type FtpConf struct {
	Path     string `yaml:"path"`
	Uri      string `yaml:"uri"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}
