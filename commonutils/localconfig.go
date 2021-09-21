package commonutils

type Config struct {
	GRPCPort int `yaml:"grpcportnum"`
	FileHost string `yaml:"filehostURL"`
	FileStorePath string `yaml:"filepath"`
}

/*
	GetConfig... returns config read from config.yaml file converting into Config struct
 */
func GetConfig() (*Config, error) {
	con := Config {
		39298,
		"http://localhost:8080/file/",
		"/tmp/",
	}
	
	return &con, nil
}
