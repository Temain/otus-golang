package configer

type ConfigApi struct {
	GrpcHost    string `config:"grpc_host,required"`
	GrpcPort    int    `config:"grpc_port,required"`
	PostgresDsn string `config:"postgres_dsn,required"`
	LogFile     string `config:"log_file,required"`
	LogLevel    string `config:"log_level,required"`
}
