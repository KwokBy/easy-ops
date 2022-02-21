package configs

// DB struct
type DB struct {
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Log      bool   `yaml:"log"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
}

// JWT struct
type JWT struct {
	Secret string `yaml:"secret"`
}
