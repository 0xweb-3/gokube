package config

type System struct {
	Addr string `yaml:"addr" json:"addr"`
	Port string `yaml:"port" json:"port"`
}

type Server struct {
	System System `yaml:"system" json:"system"`
}
